from fastapi import FastAPI, Path, Query, Body, HTTPException, Depends, APIRouter
from pydantic import BaseModel, Field, constr, conint
from typing import List, Optional, Dict, Any, Union
from fastapi.openapi.docs import get_swagger_ui_html
from fastapi.openapi.utils import get_openapi
from starlette.responses import JSONResponse
from enum import Enum
import uuid
from datetime import datetime

# Standard error response model
class ErrorResponse(BaseModel):
    """Standard error response model."""
    error: str = Field(
        ...,
        title="Error Message",
        description="Detailed error message",
        example="Invalid JSON payload"
    )
    code: int = Field(
        ...,
        title="Error Code",
        description="HTTP status code",
        example=400
    )
    details: Optional[List[str]] = Field(
        None,
        title="Error Details",
        description="Additional error details",
        example=["Field 'title' is required"]
    )

# Task models for the /tasks endpoint
class TaskCreate(BaseModel):
    """Model for creating a new task."""
    title: str = Field(
        ...,
        title="Task Title",
        description="Title of the task",
        example="Sample Task",
        min_length=1,
        max_length=100
    )
    description: Optional[str] = Field(
        None,
        title="Task Description",
        description="Detailed description of the task",
        example="Example description"
    )

class Task(BaseModel):
    """Complete task model including system-generated fields."""
    id: str = Field(
        ...,
        title="Task ID",
        description="Unique identifier for the task",
        example="f7cfc49d-824b-4728-a4c4-45e5901e3d42"
    )
    title: str = Field(
        ...,
        title="Task Title",
        description="Title of the task",
        example="Sample Task",
        min_length=1,
        max_length=100
    )
    description: Optional[str] = Field(
        None,
        title="Task Description",
        description="Detailed description of the task",
        example="Example description"
    )

# Define our product models with comprehensive documentation
class ProductCategory(str, Enum):
    """Product category enumeration."""
    ELECTRONICS = "electronics"
    CLOTHING = "clothing"
    FOOD = "food"
    BOOKS = "books"
    OTHER = "other"

class ProductBase(BaseModel):
    """Base product model with common attributes."""
    name: str = Field(..., min_length=1, max_length=100, description="Name of the product", example="Wireless Headphones")
    description: str = Field(..., description="Detailed description of the product", example="Noise-cancelling wireless headphones with 20h battery life")
    price: float = Field(..., gt=0.0, description="Price of the product in USD", example=99.99)
    category: ProductCategory = Field(..., description="Category the product belongs to", example="electronics")
    in_stock: bool = Field(default=True, description="Whether the product is in stock", example=True)
    tags: List[str] = Field(default=[], description="Tags associated with the product", example=["wireless", "audio", "bluetooth"])

class ProductCreate(ProductBase):
    """Model for creating a new product."""
    pass

class Product(ProductBase):
    """Complete product model including system fields."""
    id: str = Field(..., description="Unique identifier for the product", example="f7cfc49d-824b-4728-a4c4-45e5901e3d42")
    created_at: datetime = Field(..., description="Date and time when the product was created", example="2023-01-15T14:30:00Z")
    updated_at: datetime = Field(..., description="Date and time when the product was last updated", example="2023-01-15T14:30:00Z")

    class Config:
        from_attributes = True
        json_schema_extra = {
            "example": {
                "id": "f7cfc49d-824b-4728-a4c4-45e5901e3d42",
                "name": "Wireless Headphones",
                "description": "Noise-cancelling wireless headphones with 20h battery life",
                "price": 99.99,
                "category": "electronics",
                "in_stock": True,
                "tags": ["wireless", "audio", "bluetooth"],
                "created_at": "2023-01-15T14:30:00Z",
                "updated_at": "2023-01-15T14:30:00Z"
            }
        }


# Create the FastAPI application with detailed metadata
app = FastAPI(
    title="Product API",
    description="""
    A fully documented Product API that demonstrates documentation-driven testing.
    
    This API allows you to:
    * Create new products
    * Retrieve product information
    * Update existing products
    * Delete products
    * Search for products by various criteria
    
    ## Documentation-Driven Testing
    This API is designed for automated testing directly from the OpenAPI specification.
    Each endpoint includes complete request and response examples that can be used for:
    
    * Generating test cases
    * Validating responses
    * Performing load testing
    * Creating mock servers
    
    ## Test Automation
    The OpenAPI schema is designed to be consumed by automated test tools that can:
    1. Parse the schema to discover endpoints
    2. Extract example data for requests
    3. Validate responses against the defined schemas
    4. Generate load test scenarios
    """,
    version="1.0.0",
    openapi_version="3.0.3"  # Force OpenAPI 3.0 instead of 3.1
)

# In-memory storage for demo purposes
products_db = {}
tasks_db = {}

# Helper functions
def generate_product_id():
    return str(uuid.uuid4())

def generate_task_id():
    return str(uuid.uuid4())

def get_current_time():
    return datetime.utcnow()

# Task management routes with comprehensive documentation
@app.post(
    "/tasks",
    response_model=Task,
    status_code=201,
    summary="Create a new task",
    description="Creates a task with title and description",
    response_description="The created task with system-generated fields",
    tags=["Tasks"],
    responses={
        201: {
            "description": "Task created successfully",
            "content": {
                "application/json": {
                    "example": {
                        "id": "f7cfc49d-824b-4728-a4c4-45e5901e3d42",
                        "title": "Sample Task",
                        "description": "Example description"
                    }
                }
            }
        },
        422: {
            "description": "Validation Error",
            "content": {
                "application/json": {
                    "example": {
                        "detail": [
                            {
                                "loc": ["body", "title"],
                                "msg": "field required",
                                "type": "value_error.missing"
                            }
                        ]
                    }
                }
            }
        }
    }
)
async def create_task(
    task: TaskCreate = Body(
        ...,
        title="Task Data",
        description="The data for the new task",
        example={
            "title": "Sample Task",
            "description": "Example description"
        }
    )
):
    """
    Create a new task.
    
    Returns the complete task object including the generated ID.
    """
    task_id = generate_task_id()
    
    new_task = {
        "id": task_id,
        **task.dict()
    }
    
    tasks_db[task_id] = new_task
    return new_task

# Routes with comprehensive documentation
@app.get(
    "/products",
    response_model=List[Product],
    summary="Get all products",
    description="Retrieve a list of all products with optional filtering by category",
    response_description="A list of product objects",
    tags=["Products"],
    responses={
        200: {
            "description": "Successful response with list of products",
            "content": {
                "application/json": {
                    "examples": {
                        "all_products": {
                            "summary": "All products",
                            "value": [
                                {
                                    "id": "f7cfc49d-824b-4728-a4c4-45e5901e3d42",
                                    "name": "Wireless Headphones",
                                    "description": "Noise-cancelling wireless headphones with 20h battery life",
                                    "price": 99.99,
                                    "category": "electronics",
                                    "in_stock": True,
                                    "tags": ["wireless", "audio", "bluetooth"],
                                    "created_at": "2023-01-15T14:30:00Z",
                                    "updated_at": "2023-01-15T14:30:00Z"
                                }
                            ]
                        },
                        "filtered_products": {
                            "summary": "Filtered products by category",
                            "value": [
                                {
                                    "id": "f7cfc49d-824b-4728-a4c4-45e5901e3d42",
                                    "name": "Wireless Headphones",
                                    "description": "Noise-cancelling wireless headphones with 20h battery life",
                                    "price": 99.99,
                                    "category": "electronics",
                                    "in_stock": True,
                                    "tags": ["wireless", "audio", "bluetooth"],
                                    "created_at": "2023-01-15T14:30:00Z",
                                    "updated_at": "2023-01-15T14:30:00Z"
                                }
                            ]
                        }
                    }
                }
            }
        },
        422: {
            "description": "Validation Error",
            "content": {
                "application/json": {
                    "example": {
                        "detail": [
                            {
                                "loc": ["query", "min_price"],
                                "msg": "ensure this value is greater than 0",
                                "type": "value_error.number.not_gt"
                            }
                        ]
                    }
                }
            }
        }
    }
)
async def get_products(
    category: Optional[ProductCategory] = Query(
        None,
        title="Filter by Category",
        description="Filter products by category",
        example=ProductCategory.ELECTRONICS
    ),
    min_price: Optional[float] = Query(
        None,
        title="Minimum Price",
        description="Filter products with price greater than or equal to this value",
        example=50.0,
        ge=0
    ),
    max_price: Optional[float] = Query(
        None,
        title="Maximum Price",
        description="Filter products with price less than or equal to this value",
        example=200.0,
        ge=0
    ),
    in_stock: Optional[bool] = Query(
        None,
        title="In Stock Only",
        description="Filter products by stock availability",
        example=True
    )
):
    """
    Retrieve all products with optional filtering.
    
    This endpoint allows filtering by:
    - Category
    - Price range
    - Stock availability
    """
    filtered_products = list(products_db.values())
    
    if category:
        filtered_products = [p for p in filtered_products if p["category"] == category]
    
    if min_price is not None:
        filtered_products = [p for p in filtered_products if p["price"] >= min_price]
    
    if max_price is not None:
        filtered_products = [p for p in filtered_products if p["price"] <= max_price]
    
    if in_stock is not None:
        filtered_products = [p for p in filtered_products if p["in_stock"] == in_stock]
    
    return filtered_products

@app.get(
    "/products/{product_id}",
    response_model=Product,
    responses={
        200: {
            "description": "Product found successfully",
            "content": {
                "application/json": {
                    "example": {
                        "id": "f7cfc49d-824b-4728-a4c4-45e5901e3d42",
                        "name": "Wireless Headphones",
                        "description": "Noise-cancelling wireless headphones with 20h battery life",
                        "price": 99.99,
                        "category": "electronics",
                        "in_stock": True,
                        "tags": ["wireless", "audio", "bluetooth"],
                        "created_at": "2023-01-15T14:30:00Z",
                        "updated_at": "2023-01-15T14:30:00Z"
                    }
                }
            }
        },
        422: {
            "description": "Validation Error",
            "content": {
                "application/json": {
                    "example": {
                        "detail": [
                            {
                                "loc": ["path", "product_id"],
                                "msg": "invalid uuid format",
                                "type": "value_error.uuid"
                            }
                        ]
                    }
                }
            }
        }
    },
    summary="Get a specific product",
    description="Retrieve detailed information about a specific product by its ID",
    response_description="The requested product details",
    tags=["Products"],
)
async def get_product(
    product_id: str = Path(
        ...,
        title="Product ID",
        description="The unique identifier of the product",
        example="f7cfc49d-824b-4728-a4c4-45e5901e3d42"
    )
):
    """
    Retrieve a specific product by its ID.
    
    If the product does not exist, a 404 error is returned.
    """
    if product_id not in products_db:
        raise HTTPException(
            status_code=404, 
            detail={
                "error": "Product not found",
                "code": 404
            }
        )
    
    return products_db[product_id]

@app.post(
    "/products",
    response_model=Product,
    status_code=201,
    summary="Create a new product",
    description="Add a new product to the catalog",
    response_description="The created product with system-generated fields",
    tags=["Products"],
    responses={
        201: {
            "description": "Product created successfully",
            "content": {
                "application/json": {
                    "example": {
                        "id": "f7cfc49d-824b-4728-a4c4-45e5901e3d42",
                        "name": "Wireless Headphones",
                        "description": "Noise-cancelling wireless headphones with 20h battery life",
                        "price": 99.99,
                        "category": "electronics",
                        "in_stock": True,
                        "tags": ["wireless", "audio", "bluetooth"],
                        "created_at": "2023-01-15T14:30:00Z",
                        "updated_at": "2023-01-15T14:30:00Z"
                    }
                }
            }
        },
        422: {
            "description": "Validation Error",
            "content": {
                "application/json": {
                    "example": {
                        "detail": [
                            {
                                "loc": ["body", "price"],
                                "msg": "ensure this value is greater than 0",
                                "type": "value_error.number.not_gt"
                            }
                        ]
                    }
                }
            }
        }
    }
)
async def create_product(
    product: ProductCreate = Body(
        ...,
        title="Product Data",
        description="The data for the new product",
        example={
            "name": "Wireless Headphones",
            "description": "Noise-cancelling wireless headphones with 20h battery life",
            "price": 99.99,
            "category": "electronics",
            "in_stock": True,
            "tags": ["wireless", "audio", "bluetooth"]
        }
    )
):
    """
    Create a new product in the catalog.
    
    Returns the complete product object including the generated ID
    and timestamps.
    """
    product_id = generate_product_id()
    current_time = get_current_time()
    
    new_product = {
        "id": product_id,
        **product.dict(),
        "created_at": current_time,
        "updated_at": current_time
    }
    
    products_db[product_id] = new_product
    return new_product

@app.put(
    "/products/{product_id}",
    response_model=Product,
    responses={
        200: {
            "description": "Product updated successfully",
            "content": {
                "application/json": {
                    "example": {
                        "id": "f7cfc49d-824b-4728-a4c4-45e5901e3d42",
                        "name": "Updated Wireless Headphones",
                        "description": "Improved noise-cancelling wireless headphones with 30h battery life",
                        "price": 129.99,
                        "category": "electronics",
                        "in_stock": True,
                        "tags": ["wireless", "audio", "bluetooth", "noise-cancelling"],
                        "created_at": "2023-01-15T14:30:00Z",
                        "updated_at": "2023-01-15T14:30:00Z"
                    }
                }
            }
        },
        422: {
            "description": "Validation Error",
            "content": {
                "application/json": {
                    "example": {
                        "detail": [
                            {
                                "loc": ["body", "price"],
                                "msg": "ensure this value is greater than 0",
                                "type": "value_error.number.not_gt"
                            }
                        ]
                    }
                }
            }
        }
    },
    summary="Update a product",
    description="Update an existing product's details",
    response_description="The updated product details",
    tags=["Products"],
)
async def update_product(
    product_id: str = Path(
        ...,
        title="Product ID",
        description="The unique identifier of the product to update",
        example="f7cfc49d-824b-4728-a4c4-45e5901e3d42"
    ),
    product_update: ProductBase = Body(
        ...,
        title="Updated Product Data",
        description="The new data for the product",
        example={
            "name": "Updated Wireless Headphones",
            "description": "Improved noise-cancelling wireless headphones with 30h battery life",
            "price": 129.99,
            "category": "electronics",
            "in_stock": True,
            "tags": ["wireless", "audio", "bluetooth", "noise-cancelling"]
        }
    )
):
    """
    Update an existing product.
    
    If the product does not exist, a 404 error is returned.
    """
    if product_id not in products_db:
        raise HTTPException(
            status_code=404, 
            detail={
                "error": "Product not found",
                "code": 404
            }
        )
    
    current_product = products_db[product_id]
    updated_product = {
        **current_product,
        **product_update.dict(),
        "updated_at": get_current_time()
    }
    
    products_db[product_id] = updated_product
    return updated_product

@app.delete(
    "/products/{product_id}",
    status_code=204,
    responses={
        204: {
            "description": "Product successfully deleted",
        },
        404: {
            "description": "Product not found",
            "content": {
                "application/json": {
                    "example": {
                        "error": "Product not found",
                        "code": 404
                    }
                }
            }
        },
        422: {
            "description": "Validation Error",
            "content": {
                "application/json": {
                    "example": {
                        "detail": [
                            {
                                "loc": ["path", "product_id"],
                                "msg": "invalid uuid format",
                                "type": "value_error.uuid"
                            }
                        ]
                    }
                }
            }
        }
    },
    summary="Delete a product",
    description="Remove a product from the catalog",
    tags=["Products"],
)
async def delete_product(
    product_id: str = Path(
        ...,
        title="Product ID",
        description="The unique identifier of the product to delete",
        example="f7cfc49d-824b-4728-a4c4-45e5901e3d42"
    )
):
    """
    Delete a product from the catalog.
    
    If the product does not exist, a 404 error is returned.
    """
    if product_id not in products_db:
        raise HTTPException(
            status_code=404, 
            detail={
                "error": "Product not found",
                "code": 404
            }
        )
    
    del products_db[product_id]
    return None

# Adding some example products on startup
@app.on_event("startup")
async def startup_event():
    # Add some example products
    example_products = [
        {
            "name": "Wireless Headphones",
            "description": "Noise-cancelling wireless headphones with 20h battery life",
            "price": 99.99,
            "category": ProductCategory.ELECTRONICS,
            "in_stock": True,
            "tags": ["wireless", "audio", "bluetooth"]
        },
        {
            "name": "Cotton T-Shirt",
            "description": "Comfortable 100% cotton t-shirt, available in multiple colors",
            "price": 19.99,
            "category": ProductCategory.CLOTHING,
            "in_stock": True,
            "tags": ["cotton", "casual", "summer"]
        },
        {
            "name": "Organic Protein Bars",
            "description": "Healthy protein bars made with organic ingredients",
            "price": 24.99,
            "category": ProductCategory.FOOD,
            "in_stock": False,
            "tags": ["organic", "protein", "healthy"]
        }
    ]
    
    for product_data in example_products:
        product_id = generate_product_id()
        current_time = get_current_time()
        
        product_obj = {
            "id": product_id,
            **product_data,
            "created_at": current_time,
            "updated_at": current_time
        }
        
        products_db[product_id] = product_obj
    
    # Add some example tasks
    example_tasks = [
        {
            "title": "Sample Task",
            "description": "This is an example task"
        },
        {
            "title": "Another Task",
            "description": "This is another example task"
        }
    ]
    
    for task_data in example_tasks:
        task_id = generate_task_id()
        
        task_obj = {
            "id": task_id,
            **task_data
        }
        
        tasks_db[task_id] = task_obj

# Add special routes for test validation
@app.get(
    "/test/health",
    summary="Health check endpoint",
    description="Validate that the API is running properly",
    response_description="Health status of the API",
    tags=["Testing"],
    responses={
        200: {
            "description": "API is healthy",
            "content": {
                "application/json": {
                    "example": {
                        "status": "healthy",
                        "version": "1.0.0",
                        "timestamp": "2023-01-15T14:30:00Z"
                    }
                }
            }
        }
    }
)
async def health_check():
    """Health check endpoint for testing."""
    return {
        "status": "healthy",
        "version": "1.0.0",
        "timestamp": get_current_time().isoformat()
    }

@app.post(
    "/test/echo",
    summary="Echo test endpoint",
    description="Return the request body as-is for testing request/response handling",
    response_description="The same data that was sent in the request",
    tags=["Testing"],
    responses={
        200: {
            "description": "Echo response",
            "content": {
                "application/json": {
                    "example": {
                        "message": "This is a test message",
                        "number": 42,
                        "active": True
                    }
                }
            }
        },
        422: {
            "description": "Validation Error",
            "content": {
                "application/json": {
                    "example": {
                        "detail": [
                            {
                                "loc": ["body"],
                                "msg": "value is not a valid dict",
                                "type": "type_error.dict"
                            }
                        ]
                    }
                }
            }
        }
    }
)
async def echo_test(data: Dict[str, Any] = Body(..., example={"message": "Test message", "number": 42, "active": True})):
    """Echo the request body back as a response for testing."""
    return data

def patch_exclusive_min_max(schema):
    if isinstance(schema, dict):
        # Patch exclusiveMinimum
        if "exclusiveMinimum" in schema and isinstance(schema["exclusiveMinimum"], (int, float)):
            schema["minimum"] = schema["exclusiveMinimum"]
            schema["exclusiveMinimum"] = True
        # Patch exclusiveMaximum
        if "exclusiveMaximum" in schema and isinstance(schema["exclusiveMaximum"], (int, float)):
            schema["maximum"] = schema["exclusiveMaximum"]
            schema["exclusiveMaximum"] = True
        for v in schema.values():
            patch_exclusive_min_max(v)
    elif isinstance(schema, list):
        for item in schema:
            patch_exclusive_min_max(item)

# Custom OpenAPI schema generator
original_openapi = app.openapi

def custom_openapi():
    if app.openapi_schema:
        return app.openapi_schema
    openapi_schema = get_openapi(
        title=app.title,
        version=app.version,
        description=app.description,
        routes=app.routes,
    )
    patch_exclusive_min_max(openapi_schema)
    app.openapi_schema = openapi_schema
    return app.openapi_schema

app.openapi = custom_openapi

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8080)
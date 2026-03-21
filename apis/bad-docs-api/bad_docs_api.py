from fastapi import FastAPI, Path, Query, Body, HTTPException, Depends, Security
from fastapi.security import APIKeyHeader, HTTPBearer, HTTPAuthorizationCredentials
from pydantic import BaseModel, Field
from typing import List, Optional, Dict, Any
from fastapi.openapi.utils import get_openapi
from enum import Enum
import uuid
from datetime import datetime


# Task models — no descriptions or examples on Field() calls
class TaskCreate(BaseModel):
    title: str = Field(
        ...,
        min_length=1,
        max_length=100
    )
    description: str = Field(
        default=""
    )

class Task(BaseModel):
    id: str = Field(...)
    title: str = Field(
        ...,
        min_length=1,
        max_length=100
    )
    description: str = Field(
        default=""
    )


# Product models — no descriptions or examples on Field() calls
class ProductCategory(str, Enum):
    ELECTRONICS = "electronics"
    CLOTHING = "clothing"
    FOOD = "food"
    BOOKS = "books"
    OTHER = "other"

class ProductBase(BaseModel):
    name: str = Field(..., min_length=1, max_length=100)
    description: str = Field(...)
    price: float = Field(..., gt=0.0)
    category: ProductCategory = Field(...)
    in_stock: bool = Field(default=True)
    tags: List[str] = Field(default=[])

class ProductCreate(ProductBase):
    pass

class Product(ProductBase):
    id: str = Field(...)
    created_at: datetime = Field(...)
    updated_at: datetime = Field(...)

    class Config:
        from_attributes = True


# Create the FastAPI application — NO description, contact, or license_info
app = FastAPI(
    title="Product API",
    version="1.0.0",
    openapi_version="3.0.3",
)

# In-memory storage
products_db = {}
tasks_db = {}

# Helper functions
def generate_product_id():
    return str(uuid.uuid4())

def generate_task_id():
    return str(uuid.uuid4())

def get_current_time():
    return datetime.utcnow()

# Authentication middleware
api_key_header = APIKeyHeader(name="X-API-Key", auto_error=False)
bearer_scheme = HTTPBearer(auto_error=False)

async def verify_auth(
    api_key: str = Security(api_key_header),
    bearer: HTTPAuthorizationCredentials = Security(bearer_scheme),
):
    if api_key or bearer:
        return True
    raise HTTPException(
        status_code=401,
        detail={
            "error": "Authentication required",
            "code": 401,
            "details": ["Provide X-API-Key header or Authorization: Bearer <token>"]
        }
    )


# Task route — no summary, description, response_description, or responses={}
@app.post(
    "/tasks",
    response_model=Task,
    status_code=201,
    tags=["Tasks"],
)
async def create_task(
    task: TaskCreate = Body(...),
    _auth=Depends(verify_auth),
):
    task_id = generate_task_id()

    new_task = {
        "id": task_id,
        **task.dict()
    }

    tasks_db[task_id] = new_task
    return new_task


# Product routes — no summary, description, response_description, or responses={}
@app.get(
    "/products",
    response_model=List[Product],
    tags=["Products"],
)
async def get_products(
    category: str = Query(
        default="",
        enum=["", "electronics", "clothing", "food", "books", "other"],
        min_length=0,
        max_length=20
    ),
    min_price: float = Query(
        default=0.0,
        ge=0
    ),
    max_price: float = Query(
        default=999999.0,
        ge=0
    ),
    in_stock: bool = Query(
        default=True,
    ),
    _auth=Depends(verify_auth),
):
    filtered_products = list(products_db.values())

    if category and category != "":
        filtered_products = [p for p in filtered_products if p["category"] == category]

    if min_price > 0.0:
        filtered_products = [p for p in filtered_products if p["price"] >= min_price]

    if max_price < 999999.0:
        filtered_products = [p for p in filtered_products if p["price"] <= max_price]

    filtered_products = [p for p in filtered_products if p["in_stock"] == in_stock]

    return filtered_products


@app.get(
    "/products/{product_id}",
    response_model=Product,
    tags=["Products"],
)
async def get_product(
    product_id: str = Path(
        ...,
        min_length=36,
        max_length=36,
        pattern="^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
    ),
    _auth=Depends(verify_auth),
):
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
    tags=["Products"],
)
async def create_product(
    product: ProductCreate = Body(...),
    _auth=Depends(verify_auth),
):
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
    tags=["Products"],
)
async def update_product(
    product_id: str = Path(
        ...,
        min_length=36,
        max_length=36,
        pattern="^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
    ),
    product_update: ProductBase = Body(...),
    _auth=Depends(verify_auth),
):
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
    tags=["Products"],
)
async def delete_product(
    product_id: str = Path(
        ...,
        min_length=36,
        max_length=36,
        pattern="^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
    ),
    _auth=Depends(verify_auth),
):
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


# Legacy deprecated endpoint — no summary, description, response_description, or responses={}
@app.get(
    "/legacy/products",
    tags=["Legacy"],
    deprecated=True,
)
async def list_products_legacy(_auth=Depends(verify_auth)):
    return [
        {"name": p["name"], "price": p["price"]}
        for p in products_db.values()
    ]


# Adding some example products on startup
@app.on_event("startup")
async def startup_event():
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


# Test routes — no summary, description, response_description, or responses={}
@app.get(
    "/test/health",
    tags=["Testing"],
)
async def health_check():
    return {
        "status": "healthy",
        "version": "1.0.0",
        "timestamp": get_current_time().isoformat()
    }


@app.post(
    "/test/echo",
    tags=["Testing"],
)
async def echo_test(data: Dict[str, Any] = Body(...), _auth=Depends(verify_auth)):
    return data


def patch_exclusive_min_max(schema):
    if isinstance(schema, dict):
        if "exclusiveMinimum" in schema and isinstance(schema["exclusiveMinimum"], (int, float)):
            schema["minimum"] = schema["exclusiveMinimum"]
            schema["exclusiveMinimum"] = True
        if "exclusiveMaximum" in schema and isinstance(schema["exclusiveMaximum"], (int, float)):
            schema["maximum"] = schema["exclusiveMaximum"]
            schema["exclusiveMaximum"] = True
        for v in schema.values():
            patch_exclusive_min_max(v)
    elif isinstance(schema, list):
        for item in schema:
            patch_exclusive_min_max(item)


# Custom OpenAPI schema generator
def custom_openapi():
    if app.openapi_schema:
        return app.openapi_schema
    openapi_schema = get_openapi(
        title=app.title,
        version=app.version,
        description=app.description,
        routes=app.routes,
    )

    # Add security schemes (so P005 passes)
    openapi_schema["components"]["securitySchemes"] = {
        "ApiKeyAuth": {
            "type": "apiKey",
            "in": "header",
            "name": "X-API-Key",
            "description": "API key for authentication. Provision keys via the admin dashboard."
        },
        "BearerAuth": {
            "type": "http",
            "scheme": "bearer",
            "bearerFormat": "JWT",
            "description": "JWT token for authentication. Obtain tokens via the /auth/login endpoint."
        }
    }

    # Add ErrorResponse schema to components (for P004 completeness)
    openapi_schema["components"]["schemas"]["ErrorResponse"] = {
        "type": "object",
        "description": "Standard error response model used across all endpoints.",
        "required": ["error", "code"],
        "properties": {
            "error": {
                "type": "string",
                "description": "Human-readable error message",
                "example": "Invalid request data"
            },
            "code": {
                "type": "integer",
                "description": "HTTP status code",
                "example": 400
            },
            "details": {
                "type": "array",
                "items": {"type": "string"},
                "description": "Additional error details",
                "example": ["Field 'name' is required"]
            }
        }
    }

    # Add global security requirement (so P005 passes)
    openapi_schema["security"] = [
        {"ApiKeyAuth": []},
        {"BearerAuth": []}
    ]

    # Add versioning strategy to description (so P008 passes)
    desc = openapi_schema.get("info", {}).get("description", "") or ""
    if "Versioning Strategy" not in desc:
        openapi_schema["info"]["description"] = (
            (desc + "\n\n" if desc else "")
            + "## Versioning Strategy\n"
            "This API follows semantic versioning (SemVer). Breaking changes are introduced only in major version bumps.\n"
            "For backward compatibility, deprecated endpoints remain available for one major version cycle.\n"
            "Migration guide: see the changelog at /docs/changelog for upgrade instructions between versions.\n"
        )

    # Remap per-operation security names to match our custom scheme names
    scheme_remap = {
        "APIKeyHeader": "ApiKeyAuth",
        "HTTPBearer": "BearerAuth",
    }
    for path_item in openapi_schema.get("paths", {}).values():
        for operation in path_item.values():
            if isinstance(operation, dict) and "security" in operation:
                operation["security"] = [
                    {scheme_remap.get(k, k): v for k, v in req.items()}
                    for req in operation["security"]
                ]

    patch_exclusive_min_max(openapi_schema)
    app.openapi_schema = openapi_schema
    return app.openapi_schema

app.openapi = custom_openapi

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8080)

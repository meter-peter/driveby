# Prof. Kritikos — Inline PDF Comments (extracted)

Source: `/home/meter-peter/Downloads/thesis.pdf`
Total comments: 610

Regenerate with: `python3 extract_comments.py <pdf> <md>`


## Page 1

**[1]** (Text) Please provide the names of the two examination committee members when they become available.


## Page 7

**[2]** (Text) Missing list of images, list of tables and table of acronyms.

Each list/table should be placed on its own page.

Finally, you are missing the abstract (1-2 paragraphs with keywords at the end) in both the Greek and English language.


## Page 8

**[3]** (Text) Maybe it is better to be as precise as possible. Your focus is on RESTful APIs, not APIs in general. 
Thus, please update this in the current sentence. 

In addition, if you believe that RESTful is a word you would not like to repeat all the time, you could state that from now on, when you say API you mean RESTful API.

**[4]** (Text) Missing references to back up your claims. Please provide them as much as possible in the report.

**[5]** (Text) the development of these ...

**[6]** (Text) please provide URL of this implementation (in github) as a footnote.

**[7]** (Text) their current work

**[8]** (Text) on core development tasks, on code improvement efforts or on innovation (e.g., encompassing innovative features in the software).


## Page 9

**[9]** (Text) What is Accelerate? Maybe explain in a form of footnote.

**[10]** (Text) Could also indicate whether the advantage of automation leads to a remarkable possible impact on the organisations that feature it. And could explain what that impact is.

**[11]** (Text) an

**[12]** (Text) compress what? The development?

**[13]** (Text) So, what ...


## Page 10

**[14]** (Text) However, the

**[15]** (Text) which ones? You mean the "software systems under development and maintenance"?

**[16]** (Text) Please correct the overflow here if possible.

**[17]** (Text) Please provide a reference if this has been actually told by someone.


## Page 11

**[18]** (Text) I believe that you should not talk about your system yet. 
You still need to clarify what is the current situation and the current technologies used. 
Thus, parts of this section are proposed to be removed and potentially re-used in later places of the report.

**[19]** (Text) again we have an overflow here.

**[20]** (Text) I believe after reading the whole section that is not actually needed. So, it could be removed. 
If you need to keep it, you should only talk about technologies and what they can offer. Nothing more that that.

**[21]** (Text) in order to fulfill its main development tasks.

**[22]** (Text) and continuously improve software based on the DDT results.

**[23]** (Text) Maybe also provide a reference that proves this.


## Page 12

**[24]** (Text) This also depends on the development model and the development practices. For instance, agile methodologies that lead to shared code base can mitigate this issue ...

**[25]** (Text) So, it ..

**[26]** (Text) So, is it a common practice to include it in code? Even if it is not related to the actual code implementation (in the sense that it is not code itself)?

**[27]** (Text) I do not disagree with you here. However, someone can argue that agents can also consume this knowledge, especially if they are targeted towards it. So, it could complement an OpenAPI specification. 

Further, you do not talk about the quality of the OpenAPI specification itself. While it can represent a contract, the question is whether it can be fully understood by a machine in the sense that the semantics of all API methods/operations are crystal clear. Without proper quality and richness of API documentation, it will not be possible, for instance, to automatically derive the right tests. So, the API would not be properly or completely validated.

**[28]** (Text) a


## Page 13

**[29]** (Text) You have not talked about the nine DDT principles yet. So, please introduce them beforehand.
Otherwise, please modify the current sentence in order not to mention them.

**[30]** (Text) what about progress drift, i.e., the code that has progressed with a modification of the actual behaviour. 
The API should be also modified to reflect this.
So, the question is also whether the API must be automatically changed after permanently committed code changes.

Or do you believe that the API should be stable from the very beginning and can thus drive the whole implementation cycles?

**[31]** (Text) Again you talk about the nine principles without introducing them.

**[32]** (Text) also

**[33]** (Text) Are you sure about that? Maybe someone could argue that these are possible nowadays.


## Page 14

**[34]** (Text) I totally agree here, we have a semantic incompleteness of API specs based on the current tooling that automatically generates them.

**[35]** (Text) ok but what about machine enrichment? Maybe the machine can understand the code and provide the missing pieces?

In any case, I agree that DDT could detect in an early stage whether an API spec is sufficient for testing purposes.

**[36]** (Text) not totally agreeing here in the sense that deployment is an aftermath of testing through pipeline execution. 
Thus, there is a connection.

**[37]** (Text) There are also other tools that produce API tests from OpenAPI specification. For instance, please consider EvoMaster. So, please add sth about these efforts, too.

**[38]** (Text) correct, but all these are aspects of the OpenAPI spec itself. So, for reasoning about its semantic and syntactic correctness and completenes.
But the question is whether we need to go beyond that and also properly derive the validation of the actual sw itself that conforms to this API. 
Isn't that what you are aiming in your work? 
Conduct both API spec and code validation? And check the synchronisation between them?


## Page 15

**[39]** (Text) Maybe just provide the research questions here and then when or after presenting the thesis contributions, you could explain how they address the research questions posed. And could provide a nice mapping table from research questions to contributions.

**[40]** (Text) The three axioms are these ones? Not so clear ...

**[41]** (Text) in this thesis

**[42]** (Text) Please provide a reference for this term. Maybe also explained in a footnote.

**[43]** (Text) Not clear what these two mean here. Please explain.


## Page 16

**[44]** (Text) The acronym needs to be introduced here (or beforehand if possible and more appropriate).

**[45]** (Text) IDP as an acronym (Internal Development Platform) was never introduced before in the report. So, please introduce this acronym here (unless there is a better place beforehand).

**[46]** (Text) It is not clear what is this. Please provide reference and URL. Could shortly explain it in a footnote.

**[47]** (Text) Could also provide a table that provides a mapping between the contributions and the research questions, highlighting which contributions address (collectively) which research question(s).

**[48]** (Text) These three axioms should have been named already beforehand. Please see respective comment previously.

**[49]** (Text) Maybe indicate (either here or somewhere else in the report) that it is the first time someone attempts to formalize it.
As traditionally, it has been considered as a testing philosophy that correlates with DDD (Domain-Driven Design) and overlaps with TDD and BDD, i.e., specific formal testing methodologies.

In addition, maybe you need to attempt to "prove" that you indeed propose a formal methodology and that this is different from those already proposed and mentioned above. 
The correlation with DDD must be also made.

**[50]** (Text) please indicate where in this report are these principles specified.


## Page 17

**[51]** (Text) I am puzzled here. As you talk about the methodology that was used to automatically produce the respective CLI system that constitutes your contribution. So, how is this correlated with the consumption of DDT structured output?
This does not make so much sense in my opinion.

**[52]** (Text) Please note that a formal testing methodology should encompass not only principles and axioms but also cover various important aspects, including design, execution, interpretation and evolution along with governance.
If these are not covered, then it is considered a partial methodology.

**[53]** (Text) details ... + should you talk about the design and development of this CLI system? Or this is covered elsewhere (e.g., in Chapter 8)?

**[54]** (Text) As the architecture has not been presented in order to understand the main roles and responsibilities of each component, there is no need to provide the names of the components here.

**[55]** (Text) do you mean the adopted or implemented Kubernetes architecture?
How would you frame this under the context of one of your contributions?
Maybe indicate that the contribution X is analysed along with specific things?

**[56]** (Text) Not clear how all of these correlate with your contribution that is covered in this chapter. Maybe it is better to indicate here only what is covered in terms of your contribution and nothing more

**[57]** (Text) Again maybe correlate the chapter description with one of your contributions ...

**[58]** (Text) This is another acronym that was never defined before.


## Page 18

**[59]** (Text) this incorporates ...

**[60]** (Text) of our system

**[61]** (Text) a

**[62]** (Text) maybe provide the three main evaluations in order of strength or significance: first the proof-of-concept, then the controlled validation and finally the large-scale validation

**[63]** (Text) The AI-assisted methodology utilised for constructing the proposed system realised in the context of this thesis. 

In addition, maybe also indicate that the methodology was used for also producing the thesis report. So, the initial phrase should cover them both: "examines the AI-assisted development of the thesis system and report:" and then provide more details afterwards.

**[64]** (Text) So, DDT was utilised as an actual testing methodology for providing feedback to the agents?
This was not so clear from the previous section description - correlates with one of my previous comments.

However, one question popped out: if the methodology was used to produce the DriveBy system which automates DDT, then how is it possible to produce DriveBy through DDT? Unless you were utilised as a manual worker doing the work that DDT will eventually automate.

Please clarify in the text.

**[65]** (Text) Not clear why Chapter 9 could not be merged with Chapter 7 as I presume that Chapter 9 discusses the results from the validation/empirical evaluation that was analysed in Chapter 7. In addition, it could also include sth related to threats to validity. 

Unless you consider that Chapter 8 is a kind of additional evaluation and you need to present it before discussing overally all the results produced.

Please check and revise if necessary.

**[66]** (Text) if this is a specific contribution then there is no need to stress it separately in this sentence.
Maybe talk instead about the main benefits and impact of these contributions.

**[67]** (Text) current


## Page 19

**[68]** (Text) that justifies the need for a formal testing methodology in the form of the DDT.

**[69]** (Text) Was this categorisation adopted from [32]? If yes, this needs to be clarified here.

**[70]** (Text) not clear what local testing includes. Do we assume that service/API mocking is involved such that the ideal interactions are recorded. This would then connect well with the next part of the current sentence (verification of interactions from the provider side).


## Page 20

**[71]** (Text) Could explain what is a test stub in a footnote

**[72]** (Text) quality, completeness and synchronisation with the respective actual API implementation.

**[73]** (Text) by specification you mean the contract itself or the API specification (OpenAPI spec, for instance)?

This is not clear as a contract can be considered as a specification and by also checking the next sentence.

**[74]** (Text) So, a contract can pass, indicating that this consumer can successfully interact with the API.
But it does not cover what you mention here, which might also jeopardise the cooperation with other clients - a contract can be successful for one consumer and not successful for another (due to API incompleteness, implementation errors, lack of synchronisation between API spec and the implementation).
So, maybe you can also comment the latter here ...

**[75]** (Text) Please also add in the current analysis EvoMaster.

**[76]** (Text) Could provide a definition of this testing kind in a footnote

**[77]** (Text) what shrinking means in the current context? Not so clear ...

**[78]** (Text) fuzzing (in the context of property-based testing)

**[79]** (Text) I agree here - so, testing is covered but this does not cover the quality of the specification itself while both are important. 

Question: doesn't the quality of specification also affect the conducted tests? For instance, if RESTler does not completely understand the semantics of each API method/operation, then it cannot discover all possible method interactions or might check wrong ones. 

Thus, there is a correlation that could be made clear in the text.

**[80]** (Text) question: if a security definition is missing, wouldn't this affect the testing? If a specific endpoint requires security, its testing will fail no matter what kind of input is artificially generated and used for the testing ...


## Page 21

**[81]** (Text) ok but isn't that obvious? As this was already indicated in the previous paragraph.

**[82]** (Text) what does interactive means in practice? If we consider that the documentation is produced automatically, then how this can be interactive?
Maybe you mean that there is an interaction between the tool with a human agent in order to fill in the missing pieces in the documentation?

**[83]** (Text) I totally agree here.
A food for thought: there can be synchronisation issues between API specification and implementation. The question here is whether the API is correct and the implementation is wrong? Or that the implementation is correct and the API has to evolve in order to cover the behaviour drift? 

Please attempt to name differently both problems and indicate how existing tools could enable to cover them both.

**[84]** (Text) Ok but maybe the linters are complemented with the testing tools? Together, they can cover both aspects: documentation form and fidelity. 
So, someone could both adopt them in order to improve his/her OpenAPI specifications.

Please comment on this in the current text.

**[85]** (Text) Maybe move 2.2 as 2.1 to indicate what is the current situation with the API documentation standard before delving into details with testing and other stuff.


## Page 22

**[86]** (Text) Maybe also indicate whether the current version (3.2) is considered complete or there are specific aspects that need to be covered for the near future?
In addition, I would also add the current structure of an OpenAPI specification. What it covers and where.

**[87]** (Text) There is also OpenAPI 3.2 that has introduced improvements, such as enhanced data modeling, expanded HTTP support, and features aimed at modern API patterns (e.g., streaming). Please add this in the current paragraph. And indicate three major versions instead of two ...

**[88]** (Text) SwaggerHub might be re-mentioned here as it also provides similar facilities ...

**[89]** (Text) with respect to the code-first workflows?
I disagree here as the code-first workflows produce the specification from code. So, the specification is produced automatically in that case and can evolve as the code evolves.

**[90]** (Text) Please also check and analyze the following research prototype: 

Antonios Smardas, Kyriakos Kritikos:
Towards the Automatic Production of OpenAPI Specifications from Source Code. FiCloud 2025: 342-349

which produces rich OpenAPI specs from source code.

Still, even if such a prototype is adopted, there is the gap that the current implementation might be wrong such that the API specification becomes wrong. This is a semantic gap from intentions/design to implementation.

**[91]** (Text) I do not understand the reason for having this section as API testing was already covered in 2.1. 
Even if the tools focus on OpenAPI, these need to be mentioned and categorised based on the content of Section 2.1.


## Page 23

**[92]** (Text) is this a kind of property-based testing? How would you classify this approach?

**[93]** (Text) So, I would move this tool analysis into the section in 2.1 that talks about contract testing

**[94]** (Text) Can you please elaborate more on this to better understand it? How live API responses correlate with the mock endpoints, I do not understand this. 
Can we consider that there is a local recording of interactions with the mock endpoints and then the interactions are used to verify the real endpoints (based on issuing the same requests and matching the produced responses with the ideal ones)?

**[95]** (Text) This is the same conclusion as the one that was derived for property-based and contract testing. That is another reason for deleting the current sub-section and spreading its content in Section 2.1.

**[96]** (Text) has + provide respective reference that backs up your claim here.

**[97]** (Text) Where should the API specification converge? Or do you mean that the code should converge to the API specifications? 
Or both are needed? Not clear.


## Page 24

**[98]** (Text) I disagree here - it is too strong as an argument and wrong based on reality. 
Please check: Optic, APIClarity, 42Crunch and gateways like Kong Gateway.

Thus, please correct your statement here to be more correct and precise. 
And indicate clearly what exists and what is missing to achieve the vision.

**[99]** (Text) How can the combination of Argo CD and Argo Workflows be characterised or named? It is a complete CI/CD system supporting CI/CD, CD (deployment) and IaC?

**[100]** (Text) Can you please explain briefly these metrics and their semantics? What is considered as good performance over them?


## Page 25

**[101]** (Text) correct but there exist some governance tools like Stoplight which supply rule sets and style guides. But, of course, this is insufficient.

**[102]** (Text) OCI acronym not defined - could also explain this in a footnote

**[103]** (Text) Github-based releases

**[104]** (Text) could explain what is a container image in a footnote

**[105]** (Text) the installation via helm ...

**[106]** (Text) very good + could stress that the virtue of the quality gates, that they are API/specification-driven ...

**[107]** (Text) these things could be explained in footnotes and/or a respective reference could be given for them.

**[108]** (Text) A nice reference could even better back up your claim here.


## Page 26

**[109]** (Text) Could explain how IaC fits in the overall picture. And how it is combined with CI/CD tools like Argo CD.

Are pipelines involving IaC use to generate the infrastructure for conducting the testing of the sw as well as the infrastructure for deploying the sw once quality gates pass?

**[110]** (Text) Please note that also Terraform also supports this style of configuration/infrastructure specification.

**[111]** (Text) Thus, the ...

**[112]** (Text) This is what was also requested in my previous comment ...


## Page 27

**[113]** (Text) Maybe indicate first what is policy as code and then refer to specific tools that implement it?

**[114]** (Text) is there any kind of synergy or complementarity wrt PaC and DDT?

**[115]** (Text) reconciliation?adaptation?

Please note that I like the second term as it indicates that sth is adapted in order to reach the desired state.

**[116]** (Text) ok but in the previous instantiation of the pattern, it is assumed that the ontological specification is correct and complete. 
Here, is this guaranteed for the OpenAPI specification?
Do we consider that this specification might be written by human experts so as to be both complete and precise/correct such that it can drive the sw reconciliation?

From what I understand, this is not guaranteed. So, if your desired state has issues, how to be able to reach it through reconciliation?

Thus, I presume that there is a need for an initial validation step that drives the whole workflow only if the OpenAPI spec is appropriate. If it is, then it drives the workflow execution. If not, then nothing can happen.


## Page 28

**[117]** (Text) the agent should observe the drift between the actual and desired state and usually regards that the desired state is formally and completely described. Here this is not always possible.
Thus, there is a deviation from the pattern expected principles or axioms.

**[118]** (Text) ok - in this case, there is no reconciliation. Only the gap is detected and reconciliation could occur in both possible directions via human or LLM/AI-based agents (depending also on the drift/gap specification / report).

**[119]** (Text) maybe also precision matters here, i.e., that the observations are correct, they reflect the reality. 
Sth similar to flaky tests, for example. If a test is flaky, it needs to be corrected so as not to negatively impact the testing and development process. Similarly, the observations should be correct and consistent. The same correct observations should be supplied for the same situation/context.

**[120]** (Text) again these must be correct/precise.

**[121]** (Text) i.e., the OpenAPI specification's ...

**[122]** (Text) ok but can they guarantee the correction?
Can the correct be deterministic or this does not matter? So long that the correction leads to the desired state?
Please clarify this in the text.

**[123]** (Text) that

**[124]** (Text) (which represents the testing report and includes the actionable signals)?

**[125]** (Text) , for example, ...

**[126]** (Text) so the focus is on updating the OpenAPI specification, the code or both?

The last sentence seems to promote the OpenAPI specification improvement. 

Previous sentences are more general and might also cover correcting the code. 

Please clarify what actually concerns your system: does it cover these both?


## Page 29

**[127]** (Text) Apart from hiding cloud-native complexity, what are the additional benefits?

**[128]** (Text) ok but what is the benefit from this structuring? What is the added-value?

**[129]** (Text) of what? The IDP and the self-service capabilities?

**[130]** (Text) how and for what purpose?

**[131]** (Text) what kind of capabilities are meant here? Business capabilities encompassed in sw like RESTful APIs?

**[132]** (Text) ok but a platform consumer is different from the API consumer, right? An API consumer can be a client or a specific organisational team/department. The platform consumer is a consumer of a platform like an IDP one. 
So, these are two different things, right?

**[133]** (Text) So, 6 principles map directly to the OpenAPI specification and its quality/completeness?

What is the focus of the other three (nine in total)?

**[134]** (Text) correct, if the focus is on the specification.
If the focus is on the code, is this covered?
The code is another artifact so it can also be quality-checked. Thus, are there principles focusing on the code?


## Page 30

**[135]** (Text) are these supplied by humans? Are there standards that can be followed for them? E.g., OpenAPI for instance?

**[136]** (Text) Is there any reference that formally describes this protocol? 

There is a need to be as formal as possible in the report.

**[137]** (Text) ok but is this complementary with respect to the respective files that need to be in place?

I ask this as the DDT feedback is about correcting? So, it should be complementary to the main principles and design objectives that guide the system construction. 
In simple words: an OpenAPI could be the main guidance file while testing output can be the correction feedback.


## Page 31

**[138]** (Text) again the focus seems to be on the specifications? But the code is equally important.

**[139]** (Text) could explain how these were derived. A respective reference could make this selection (of dimensions) even stronger.

**[140]** (Text) Ok but there can be an overlap between spec quality assessment and static analysis. It could be regarded that static analysis can facilitate both: syntactic validation and spec quality. So, maybe replace static analysis with Syntactic Validation? It makes more sense and separates the target (e.g., syntactic validation) from the way it is reached (e.g., static analysis).

**[141]** (Text) please provide a reference that backs up this claim.


## Page 32

**[142]** (Text) Method. -> mapping to methodology

**[143]** (Text) Someone can understand what tick and X signify in the table. However, it is not clear what is the semantics of ~ and -. In fact, '-' could be regarded as synonymous to 'X'. Unless their difference is not clarified. 

So, please indicate either at the bottom of the table or at the respective main text what is the semantics of all 4 symbols.

**[144]** (Text) ok but isn't the contract yet another specification? Couldn't the contract also include the OpenAPI specification as part of it (e.g., indicating that this specification should be respected by the API provider)?
Someone could even argue that OpenAPI-driven testing is contract-first. 

But, of course, this depends on what is meant by contract.

Here you mean consumer-contract-first and not contract-first. In the literature, contract-first is synonymous to specification-first in the sense that the contract is a formal interface/API specification. 

In order to avoid overloading the term contract, I propose to replace contract-first with consumer-first in the sense that the consumer supplies the "contract". 
This will handle properly the terminology issue. 
But it has to be applied across the whole report (as contact-first as a term is utilised in multiple places).

**[145]** (Text) This is the contract-first, API-first or specification-first approach! Most well-established term is contract-first. So, could use that one. 
On the other hand, specification-first, while less established, signifies the existence of a normative document/specification as the source of truth. So, it is more close to your vision. 
It also implies that even a more complete OpenAPI specification is not enough and more semantics is needed.

**[146]** (Text) ok - so call it specification-first to be more close to this term!

**[147]** (Text) So, it is more natural to rely on it, right? Could indicate that in the sentence.


## Page 33

**[148]** (Text) research?

**[149]** (Text) ok - but were all of these gaps more or less covered in Section 2.6? Or propagate from conclusions drawn in other previous sections of Chapter 2? Please clarify that.

**[150]** (Text) complete?informed? Please add the right term/adjective?

**[151]** (Text) and semantically-rich?

**[152]** (Text) what is this? Please clarify in footnote

**[153]** (Text) Ok but based on the analysis, it seems that this gap overlaps with quality assessment of OpenAPI specs. If the spec is complete, it is knowledge-transferable.


## Page 34

**[154]** (Text) why not other principles (from the 9) that target the OpenAPI specification (like security clarity)?
I presume that these are also important and must be satisfied.

**[155]** (Text) Gap 3 -> always use capital for the word (like image, table, section, etc.)

**[156]** (Text) ok but how this relates to another gap? As you establish gap correlations here.

**[157]** (Text) See previous comment ...
It seems that you go beyond correlations to justify the existence for some gaps. This needs to be also communicated to the reader at the beginning of the paragraph. Or try to separate these two things (correlation vs justification).

**[158]** (Text) ok but also relates to gap 4?


## Page 35

**[159]** (Text) The Documentation-Driven Methodology 

++ Maybe call it Specification-Driven Testing (SDT) Methodology?
 As I suggested to replace documentation-driven with specification-driven or even specification-first (so specification-first testing (SFT) is another candidate here).
Please check which candidate is better in terms of your beliefs and preferences (SDT or SFT). Maybe SFT is better?

**[160]** (Text) Actually what you are defining is a conceptual framework for a formal methodology (or a strong theoretical core for such a methodology). There are various other elements that are missing, including derivation procedure for tests, execution & failure semantics and lifecycle governance. 

What you specify looks more close to a formal validation theory ...

So, maybe you need to modify the wording in order to more properly frame your contribution across the whole report!

**[161]** (Text) ok but I am a little bit puzzled as here you say "complete" while the principles that you define attempt to validate this. So, there is an inconsistency about what is actually tested. It is the code or the OpenAPI specification itself?

**[162]** (Text) validation results?

**[163]** (Text) As indicated in a previous comment, I would also add preciseness or accuracy as another property or axiom. It is different from completeness in my opinion and has to be added.


## Page 36

**[164]** (Text) in the context of a specific principle?

I propose to add this as I presume that checks stem from the principles, which practically realize the axioms as you mentioned above.

**[165]** (Text) Food for thought: does the OpenAPI cover all possible information? In order to cover both the semantics and behaviour of the API?
For instance, how can business rules be encoded (e.g., that one param value must be greater than another)?

Thus, even OpenAPI standard could be regarded as insufficient for producing all possible validation!
Someone needs to exploit its extension mechanisms in order to cover the missing parts!!!

**[166]** (Text) I would also add parameter constraints (unless you regard that such constraints are natural part of the parameter specification).

**[167]** (Text) and is important in the context of API selection & usage/integration?

**[168]** (Text) interpretation? Might be better than "reading".


## Page 37

**[169]** (Text) still this does not check whether the incorporated information is correct, i.e., it reflects the actual reality. And it might still be insufficient based on I indicated in a previous comment.

**[170]** (Text) do you mean the same API implementation? 
In the sense that one specification can be implemented and respected by different API implementations. 
The specification is a kind of a class while the API implementations its instances.

**[171]** (Text) why isn't non-functional testing also covered? As this depends on the context?

**[172]** (Text) Question: if an API produces different responses/outputs for the same test in different contexts, does this mean that the test is context-dependent and must be eliminated?
What happens with context-aware APIs? As this could be considered as a specialised case that might break determinism.

**[173]** (Text) I agree here. This relates to the concept of flaky tests. Maybe this could be commented here in the text. 

So, the path is specification -> precise & consistent tests -> deterministic output.


## Page 38

**[174]** (Text) as you state axioms as requirements, you could slightly change the wording.

E.g., DDT must produce reports ...

**[175]** (Text) this might be however verbose - the others might not be (0/1, class or check result).

Unless the remediations are taken from a remediation taxonomy.
Still it might not be clear to which element the remediation must be applied. 

This looks like a problem of formality. You need a machine-processable format that can be understood by anyone - like an ontological specification!!!

**[176]** (Text) ok but the functional testing does not contribute to observability? 

Someone would assume that it does as it can influence the remediation. 
Why a performance evaluation result might not influence directly the actual code/sw but the infrastructure on which it is executed/hosted. 
Thus, it can be argued that functional tests are more close to remediations than performance tests!

Please clarify and elaborate more in the text.

**[177]** (Text) ok but the more kinds of actors are involved, the more are the expectations and more information might have to be covered.
Further, it is of great concern if the information, even if specified in machine-processable format like XML, can be easily interpreted, even by automated agents.
That are very important aspects, which while in principle are valid, are very hard to check and achieve.

**[178]** (Text) and analysis?

As later on you talk about trends ...

**[179]** (Text) what is meant by "process" here? Please clarify in the text.

**[180]** (Text) ok but before analysing the axioms, you should explain the semantics of the severity classes. What is meant by critical, warning, etc.

**[181]** (Text) , which have complete theoretical definitions but their implementations are pending

**[182]** (Text) are these standard checks? Are they applied by linters, for instance?
You need to explain how were they derived. They look like syntax validation rules that might stem from linters.


## Page 39

**[183]** (Text) according to the designated version?
As it can be the case that a specification is said to be 3.1 but it is actually 3.2. So, someone forgot to update the version ...

**[184]** (Text) not clear what is the difference. Can you clarify this?

**[185]** (Text) the explanations aren't the same as the description? 
Maybe it is better to indicate that parameters come along with constraints on them?

**[186]** (Text) examples can also cover parameters, especially when their expected value is not clear (e.g., when their type is just a String).

**[187]** (Text) where? Maybe you should indicate where this information can be found.
This must be indicated for all the checks for this axiom.

**[188]** (Text) Please again note that this is more about information richness at the structural level. There are checks which attempt to see whether there is information richness in specific elements of the specification. This is still structural and not semantic. 
Semantic would mean that: (a) the information is precise and (b) the information is formally specified (e.g., in ontologies or logical expressions) such that it can be understood by machines.

**[189]** (Text) and example values?

**[190]** (Text) what does this mean? Please clarify in the text (or at least in a footnote).


## Page 40

**[191]** (Text) who enforces this mode? Can this be a configuration parameter? Please clarify in the text.

**[192]** (Text) there is also a test-ready mode. I do not see sth about it. So, this means that potentially either minimal or strict mode is equivalent to the test-ready mode in the current principle.

This needs to be clarified in every principle that is mode-specific and not all modes are covered.

**[193]** (Text) I agree here. Still what we see is structural. So, it does not convey to the semantic level. 
For instance, do we see 400 errors on all read & post & update operations? Shouldn't resource-specific paths have 404 errors?
This requires a deeper understanding of the RESTful semantics and in some cases of the designated operation/method. That is why they can be considered as more semantic checks ...

**[194]** (Text) Is this a matter of completeness or heterogeneity? As it relates to the way errors are formatted. So, some errors would follow format 1 and others format 2. Both formats can be complete as they convey all the necessary information. Thus, it is a matter of format uniformity rather than error specification /format completeness.

This raises the question: is uniformity a distinct property or axiom that could be taken into account?
It could also spread into other subjects like path uniformity, which is not checked.

In my opinion, the question is whether the API specification quality as a whole must be checked or only its completeness. 
As quality can influence not only testing (e.g., an agent might confuse the naming schemes applied in the spec and produce wrong tests) but also the very selection and re-use of an API. 

Thus, in my opinion, it is worth being covered but the question is whether this coverage should be totally complete or sufficiently complete to cover mainly the testing goal. 

In any case, I would separate uniformity checks from completeness ones as they control a different property.

**[195]** (Text) As some principles differ in terms of their detailed implementation based on the mode, this needs to be clarified at the metadata level in the description of each principle. 
Thus, could have a third property named as mode-specific behaviour indicating whether the tests produced and executed for the principle are different in different modes.

**[196]** (Text) structural and informative?
Structural is needed in order to not stay at a superficial high-level. But completely cover the whole structure of a schema. 
Informative is more about the validation of values in the components/instances (requests/responses) of the schema.

**[197]** (Text) As in the case of security levels, the modes need also to be formally defined before delving into the principle details.

**[198]** (Text) for them.

**[199]** (Text) Question: I presume that severity means ability to test the API. In this context, if types are missing, isn't this of a critical severity?
How can someone test an API, if it does not know the types of the parameters in its methods?

Thus, I see a potential overlap between modes, completeness and severity.
Without seeing the semantics of severity and modes, I am not sure whether this is a real overlap and also its degree. 

IN any case, please consider my comment and see whether there is a need to fix things.

**[200]** (Text) ok but for the representation schemas, shouldn't we also check the schema structure? \tWhether it exists and includes specific properties?
Or this could depend on the respective case?
In principle, both request and response bodies should have structured schemas. This highlights/signifies the quality also of the API.

**[201]** (Text) Maybe call this value schema while content schema can be called representation schema.
This is more standard terminology.

**[202]** (Text) shouldn't path parameters be required?

**[203]** (Text) ok but these are not checked in terms of their presence (i.e., a representation schema certainly has object properties).


## Page 41

**[204]** (Text) why this can be considered as minimal? Can you justify this?

**[205]** (Text) This check needs proper justification. Why these two schemes are needed? Are there other alternative schemes that could be used instead?

**[206]** (Text) Maybe this check should also be done for the previous principle?
It can be possible that one custom type/schema is missing while mentioned as a type of request or response body.
The same for schema components. A schema could have an undefined reference sub-schema.
Please add these checks in the previous principle.

**[207]** (Text) You provide some implementation details. I am not sure these are relevant at this point.
For instance, this could raise the question what is this helper and in what form it has been implemented and which checks does it cover (one helper could cover multiple and not just one check).

**[208]** (Text) I propose to have this status in all principle descriptions, not just the current one.

So, the metadata should include axiom, severity, mode-specific behaviour (yes/no) and status.

**[209]** (Text) is this an interface implemented by every principle checker? How each checker is implemented? How helpers relate to checkers? 
As you can see, reference to implementation terms/details raises questions at this point and could be avoided.


## Page 42

**[210]** (Text) ok but what is the actual difficulty? As I presume that the framework already has other implementations of principles that have been integrated and the current issue looks like a structural/signature one (right?)
Or is it a matter of what should be written in the report?
Thus, only this aspect has not been covered yet?

**[211]** (Text) ok but how was this established?
As it could be regarded as difficult.
Maybe this relates to the need to chain method calls together?
So, it is related to other works?
Please clarify slightly with a reference to these works and then detail in another chapter.

**[212]** (Text) ok but are these performance expectations imprinted in the OpenAPI spec? If yes, what is the place where they are situated?
I have the feeling that this standard does not cover them ...
So, there is a need to rely on an extension to cover them.

**[213]** (Text) please note that metrics need to be precisely defined or at least named such that it is clear how they can be measured. 

For instance, average availability can mean different things depending on how it is actually measured. 
Thus, the proper way to cover performance requirements is to "invent" an extension to OpenAPI and have a specific description of the requirements that includes the formal specification of the respective metrics.

I presume that this could be covered by future work (in this complete form) ...

**[214]** (Text) but are these part of the OpenAPI specification or not? As this was regarded as the sole artifact on which the tests are generated.

**[215]** (Text) But interestingly, the performance testing is context-aware. So, I am not sure that the same results can be produced for the same spec and API impl. as the results can depend on the current context (demand, hw configuration, etc.). 
Thus, performance testing has the potential and risk to violate the determinism constraint, even if it contributes to observability. 

Question: if it is about requirements checking / conformance, then why isn't related to determinism? As determinism is about checking requirements satisfaction: it should be 0 or 1 whether a requirement is met or not. And this should not depend on the context as you state. 
But this also relates to the way the requirement is expressed - if it is tied to a respective contextual constraint, then it could make sense deterministically (e.g., if you deploy the API in a resource with at least these capabilities, then the expected performance must be higher than this level ...)

**[216]** (Text) does this justify the classification or correlation of the principle with the observability axion? This correlation should be properly justified.

**[217]** (Text) again, why is this considered minimal? Maybe just the existence of version is minimal?

Of course, I agree with you. I just indicate that justification should be given for the choices made.


## Page 43

**[218]** (Text) NO, it is not the API version. It can mean the version of the API specification. The versioning depends on the strategy applied by an organisation. So, it can be evidenced in different places of the OpenAPI specification.

The question is whether your checks cover one or any possible versioning strategy and are independent of any relevant policies of a specific organisation.

Both of these have to be guaranteed by the checks!

**[219]** (Text) isn't this implementation detail?

**[220]** (Text) maybe this is incomplete as version could be also a domain-specific word in some cases ...

**[221]** (Text) I presume that this is the minimal requirement. The other requirement looks like somehow team/organisation-specific and not generic one. Can you please justify it? For instance, it might be proper in the description to indicate which method to use instead without requiring to use a word like "deprecat".
Please also note that this is not a complete word ...

**[222]** (Text) It is not clear how these are covered through checks. We understand the semantics of the checks but not how they are realised. If this is covered in a later chapter, it is ok. 
If not, it must be done here or there.

**[223]** (Text) So, this is more related to governance practices, right?
The main issue here is to guarantee that governance is independent from any organisation. I hope this is guaranteed by your check but it has to be well justified in the text (potentially here and not elsewhere).

**[224]** (Text) This overlaps with other principles. I do not understand why it has been specified as a distinct one (also due to this overlap).
I had the impression that test readiness is actually guaranteed by many other principles, especially those related to completeness!

**[225]** (Text) I am not convinced about these checks. They are already covered in other principles or the checks in these principles could be slightly extended to cover them (only part of the check has to be added).


## Page 44

**[226]** (Text) the precondition isn't the completeness of the specification?
Thus, to me, this classification is also wrong as the checks naturally map to completeness ...

**[227]** (Text) if deterministic test execution is not guaranteed without this principle, then why its severity is warning?

**[228]** (Text) I do not see any additional detail in figure that is not covered by the table. Usually, it is not required to have two graphical elements that greatly overlap with each other. One that covers all of the information is enough.
So, I would propose to remove the figure.

**[229]** (Text) What about P005. It looks to me that it is also related to completeness.

**[230]** (Text) so they also relate to determinism?
I ask this as you correlate them only with completeness and not determinism.

While definitely they affect determinism. 
Maybe distinguish between the very nature of a principle from its impact. 
So, nature for P001-P005 is completeness and impact is on determinism as the checks are deterministic and they also enable determinism for other checks (e.g., functional testing ones).


## Page 45

**[231]** (Text) This should be critical

**[232]** (Text) Not so clear why P009 is really difficult or non-trivial.
The checks seem easy to implement. The impact is clear as examples can be re-used for functional test execution so they could in principle eliminate fuzziness. So, what is the actual problem with them?

**[233]** (Text) Ok but in my opinion, this is a doubly-natured principle in the sense that it relates to completeness (is specification complete) and misconfiguration (related to observability).

**[234]** (Text) If an API versioning strategy is not backwards compatible, then we have immediate drastic disruption with breaking changes at the client side. However, quality-wise indeed they do not indicate that the (server-based/API) system can fail. 
Concerning performance, I am not actually convinced as potentially a badly implemented API could immediately fail even with the smallest possible load. Thus, we can have catastrophic failure.
Further, at the business side, failure to comply with non-functional requirement can be catastrophic as you lose clients and performance might be inadequate for the application domain. 
Thus, I am not sure whether performance is really "warning". If we also put in the table conformance checking, things can get even worse.

**[235]** (Text) Very nice - but the definition of the semantics has to be moved earlier in the Section 3.3


## Page 46

**[236]** (Text) This could be also moved earlier

**[237]** (Text) What I understand is that all principles are always evaluated. The mode influence the evaluation strength or degree. The stricter is the mode, the more intensive/detailed evaluation work is conducted.
Thus, I tend to disagree with your point here (although I can understand that reflect the way your system has been implemented).

**[238]** (Text) correct but the validation can be minimal in any principle. It does not have to lead to selecting specific principles only.
This looks counter-intuitive in my opinion.

**[239]** (Text) In the description of 3 principles (versioning, security and error handling), you have indicated that the minimal mode applies. Here we only see P001 being applied while this was not evident in its description.
This is a major inconsistency that has to be corrected!!!

**[240]** (Text) Please attempt to make this description fully-justified. Same for following.

**[241]** (Text) but they are minimally based on their descriptions/analysis! 

Please check the different modes and be consistent everywhere (table, current text, analysis text of each principle).

**[242]** (Text) I do not really understand this mode. The other principles have to be satisfied in order for the mode to be successful. So, it is not clear why they are neglected. 
Maybe there is a specific assumption about the selection of each mode.
Here we can assume that the API specification is rather complete and guarantees deterministic test execution? This would then make sense to focus on real testing ...

OK - it is explained later on ...
But please provide the explanation first and then clarify what is involved in each mode. This should be better way to present them.


## Page 47

**[243]** (Text) are you sure that P002,  P003 and P005 are not relevant? I have the impression that they are. 
Maybe this correlates with the overlap between P009 and other principles. But I am not sure whether P009 covers well what the other overlapped principles cover ...

**[244]** (Text) a release shouldn't be checked against the real tests, functional and performance? I do not get it.

**[245]** (Text) I see mainly minimal -> strict -> test-only as more meaningful stages/phases, which also seem to be somehow incremental (the first two are incremental based on tests, but the sequence can be considered as incremental also as it indicates the actual progression towards the real tests ...). 
The severity or intense of testing is also incremental (we expect more tests to be node on second phase and even more on the last one).


## Page 48

**[246]** (Text) The DriveBy System. 

I prefer this title. The CLI Architecture indicates that you are touching generally the architecture of a CLI, which is not your actual intention.

**[247]** (Text) Maybe you also need to explain the overall development process, such as which were the main requirements to satisfy, whether particular UML diagrams were designed and which implementation technologies were utilised and how they connect with each other.

**[248]** (Text) The presentation only covers the data flow or also the overall system/tool architecture?

**[249]** (Text) system

"binary" is an implementation-specific term"

**[250]** (Text) can we consider these as high-level requirements that drove the design and implementation of your system?

**[251]** (Text) Ok but there are structural differences between the different versions while new features/elements are introduced (in newer versions).
So, I do not understand how is it possible to not have version-specific branching logic.

On the contrary, what I understand is that the main goals of checking within the principles remain the same. This is what is invariable.
The checking logic itself is necessarily variable.

**[252]** (Text) you need to provide the design goals in an architecture-agnostic manner.
So, reference to specific architectural elements should not be supplied unless the main functionality of such elements is obvious but also explained.
For instance here, the engine is what? The whole tool? A specific component of the tool that does what?


## Page 49

**[253]** (Text) based on the considered principles, this is an almost valid design goal.
The sole exception is the principle about performance testing where the respective results have the tendency to vary.

**[254]** (Text) apart from the structuring, is there any other kind of related need? 
For instance, the content should not be subject to different interpretations.
I presume that this was mentioned as a requirement in the previous chapter in the context of the observability axiom.

**[255]** (Text) in principle - the content may vary.

**[256]** (Text) only these two were applied? Any other?

**[257]** (Text) ok - so there is a common interface for checkers that is adopted independently of the OpenAPI version.
This is what you indicate here.

**[258]** (Text) overflow here ...

**[259]** (Text) you forgot to say sth here about two other packages: testing and loader.


## Page 50

**[260]** (Text) could also indicate some words about each layer.
For instance, what is generally offered by the interface layer and why does it have two levels. Is the first one user-oriented and the second backend-oriented?

**[261]** (Text) Question here: in terms of dependency, it seems that the core tests (functional & performance) should be done in the context of specific principles. However, these seems to be in a higher level. This is not so clear and requires some justification.

**[262]** (Text) Nine checkers you have designed, right? Not 8 ones.


## Page 51

**[263]** (Text) The last three packages were not mentioned before. Can we consider them as orthogonal? Or can we consider that they reside at an infrastructural layer?
Please clarify in the main text.

**[264]** (Text) Suddenly, you provide here some code without explaining what was the implementation language used!

Further, it is ok if you supply some code fragments for demonstration purposes (the current one is ok) but do not overdo it. In case you need to provide extensive parts of code, this should be done in the Appendix.

**[265]** (Text) very minor: you have a comment extending (visually) one like while it is a one-line comment.


## Page 52

**[266]** (Text) Question here: testmode was not introduced in the previous chapter. What is its current semantics?

Similarly, while test status is more or less understandable, there are some values/members of it that require explanation like incomplete. What incomplete means in practice? In addition, why a specific test can be skipped? Is this related to the validation mode or other factors?


## Page 53

**[267]** (Text) Ok, here you indicate the overall validation result for a specific principle. What about the respective checks? Are these covered and where? The developer or agent needs to know which tests passed and which not and might be also interested to see the possible remediation.

**[268]** (Text) in general? Or in the context of specific tests?
It could be possible to have categories of tests being impacted.
For instance, if error msgs are not described, then "negative" functional tests are impacted (they cannot be executed).

**[269]** (Text) Ok - maybe refer to the Appendix to see the definition of this struct?

**[270]** (Text) ok but the checks are not well mapped to the principle results in my opinion. Shouldn't we know which checks passed or not?


## Page 54

**[271]** (Text) Correct - I totally agree here!

**[272]** (Text) overflow

**[273]** (Text) very good. If a new feature is involved in the OpenAPI standard that requires the creation of a new checker, would that mean that the APISpec interface would need to be modified/extended? How could this affect the implementing adapters?
Please attempt to comment in the text.

**[274]** (Text) Table overflows significantly. Please correct - the third column content is cut and cannot be completely seen!

**[275]** (Text) please explain in footnote what this means.

**[276]** (Text) how? Can you provide a respective example?


## Page 55

**[277]** (Text) alternatively

as one of them will be used in the context of a specific OpenAPI spec.

**[278]** (Text) Could explain in text at high-level the main logic of this code.


## Page 56

**[279]** (Text) ok - but maybe there could be another way to solve this problem?
As sanity checks do not support full validation, I presume.
If there was a different library used for this purpose?

**[280]** (Text) what is spec.go? Is the implementation of the ApiSpec?


## Page 57

**[281]** (Text) overflow again here


## Page 58

**[282]** (Text) overflow again here.

**[283]** (Text) so they downgrade into 3.0.x some parts of the document in order to enable its validation. 
Please indicate that - very clever ...

**[284]** (Text) Maybe the justification for adopting this library could be mentioned beforehand so as to know exactly the main reasons for this choice

**[285]** (Text) Question: a checker implements all tests as presented in previous chapter?
I presume yes. 
So, it is a kind of testing class for the whole principle.

**[286]** (Text) what does this mean? Maybe explain in footnote.
Please note that I have seen context also in previous listings.


## Page 59

**[287]** (Text) question: this method covers both version cases for OpenAPI? So, it also conducts sanity checks in case of OpenAPI 2.x?


## Page 60

**[288]** (Text) Overflow also here

**[289]** (Text) Please "check" whether the number of checks is correct per principle.

**[290]** (Text) ok - could also justify why P003 is the second largest checker. 
And why P008 is the smallest.

**[291]** (Text) Here you provide only the implemented ones, right?


## Page 61

**[292]** (Text) Not implemented yet ...

**[293]** (Text) what is meant by this? What is a combined mode? Do you mean that someone could select multiple modes? Does this make sense? Please clarify

**[294]** (Text) so, it is a kind of backend? That interfaces with the CLI (front-end)?


## Page 62

**[295]** (Text) do we see it somewhere?
Could be shown in the Appendix.


## Page 63

**[296]** (Text) ok but this has to be proven. Maybe point to the evaluation chapter here for this purpose.

**[297]** (Text) please note that this covers single interactions.
However, in some cases, composite interactions might need to take place. Further, prerequisites (e.g., DB state) are not covered at all.
Can we consider that your work here is rather preliminary and will be enriched in the near future to become more complete?

**[298]** (Text) how does it find if an operation is deprecated or not?


## Page 64

**[299]** (Text) ok but this looks like fallback, it might fail to provide a correct test.

**[300]** (Text) How are these supplied? Please clarify in the text.

**[301]** (Text) ok but this looks like blind validation in the sense that it is not clear what should have been the correct response.
For instance, if you receive a 4XX error, what does this mean? That the request was indeed invalid (based on the spec) or that there is a change in the implementation that causes the request to be invalid (while it is apparently valid based on the OpenAPI spec).

There is a need to have an oracle to be able to conduct properly the tests.

Alternatively, you could exploit an existing API testing approach like the ones mentioned in a previous chapter.
It is not clear why did you try to produce your own API testing approach. 
What were the main issues in the literature that you intended to solve?

**[302]** (Text) So, skipping indicates synchronisation issues between spec & implementation in the context of specific aspects, like the security one?

**[303]** (Text) "reporting ...": the API reports the requirement in its responses? Please clarify that in the text.

**[304]** (Text) so, maybe the other testing approaches are not deterministic?
That is why you implemented your own?

**[305]** (Text) ok but this is one kind of performance testing and many others exist.
I presume that this is just the beginning and more will come in the future.

**[306]** (Text) What does this mean? Please clarify and also indicate how this is achieved.

**[307]** (Text) I do not like the term "attack" here. You do not conduct penetration/security testing but load testing. 

In addition, in load testing, shouldn't you also increase the load? Does rate covers this? It is not clear what is meas. Please clarify.

**[308]** (Text) Ok but how are these defined? Are they part of the spec? Or they are part of the pipeline?

**[309]** (Text) Ok but this was already mentioned above for both of these testing kinds. No need to repeat it.


## Page 65

**[310]** (Text) So, this is "static" information coming from configuration files or command flags?

**[311]** (Text) question: the version or deployment is not covered in the name of the report. So, someone needs to consult another source to find it. Is this ok? which is the other source to consult?

**[312]** (Text) So, this is not done by the JSON renderer?
Or this is done by both of them?

**[313]** (Text) test-kind-specific ...

**[314]** (Text) Ok - but maybe a schema could be good to have here? As it gives predictability: you know exactly what you consume and where to find the information that you are interested in.

**[315]** (Text) maybe for now. But in the future, it might also incorporate them - e.g., they could also facilitate some checks, like the functional testing and versioning principle ones.

**[316]** (Text) ok - so, we always talk about one remediation. Thus, there are not multiple alternative remediations?
If yes, then the agent should choose which one to select.


## Page 66

**[317]** (Text) two overflows exist in this page

**[318]** (Text) ok - but maybe also indicate what is the file for implementing the CLI itself?

Maybe also a package diagram and "class" diagram could show everything (in terms of your implementation)?


## Page 67

**[319]** (Text) if the spec is missing, is this an infrastructure problem? Even if you retry, the problem will remain. In contrast to a transient network failure (which is indeed a problem of infrastructure).

**[320]** (Text) ok but we did not see architecture and data flow diagrams.
Only a description of packages and their most important parts.
In my opinion, this is insufficient and should be corrected/improved.


## Page 68

**[321]** (Text) question: this verification requires the API to be somewhere available.
Should we consider that this could be locally (e.g., where the system also compiles and runs the respective code) or the user supplies a URL of an already deployed API?


## Page 69

**[322]** (Text) Please remove "of Chapter 2". It is obvious.


## Page 70

**[323]** (Text) ok but how is this achieved?Do you have different binary versions per each OS?

Or this relates to Stage 3 so it is not absolutely true (for the current stage)?

**[324]** (Text) are these two involved in the next stage and the last one is about the final stage?
As we have three things and two next stages.


## Page 71

**[325]** (Text) he/she declares

**[326]** (Text) no matter how ...

**[327]** (Text) a


## Page 72

**[328]** (Text) It is not clear what are these abstraction layers. Can you depict a nice figure that shows them?
Please note that previously you talked about stages and not layers.

**[329]** (Text) where is this layer placed? It is the highest possible?

**[330]** (Text) Again, this table overflows.
Please fix this

**[331]** (Text) was there something implementation-specific that must be highlighted here? Or this is covered in the next sections?

Maybe also supply a nice image that shows this layer (what are its levels and components)?

**[332]** (Text) ok but from where do these come from?
You need to clarify how you derive your goals or requirements that drove the system development.


## Page 73

**[333]** (Text) how these correlate with the validation logic that was implemented in the CLI? It is not clear to the reader. Maybe explain it here or in footnote.

**[334]** (Text) can you please elaborate more on this? Which is also higher than the other?

**[335]** (Text) (developed by the platform team)

**[336]** (Text) What does system context mean?
Is this a kind of context diagram that you are providing or something different?


## Page 74

**[337]** (Text) It is also not clear in this figure what is already there and what was implemented in the context of the thesis.
Maybe this is clarified?

I propose to explain this figure as this could also showcase the needed clarifications that I request above.

**[338]** (Text) do the colours play a role in the diagram? If yes, please indicate their semantics

**[339]** (Text) ok, I agree. But the key is the state. It has to be permanent and transactional, right?


## Page 75

**[340]** (Text) is this a component of a Kubernetes controller or the controller itself?
As you slightly modified the terminology here.

**[341]** (Text) gradually

**[342]** (Text) an


## Page 76

**[343]** (Text) Ok but these are provider-specific for the Github provider. So, here you supply a specific example of how a composition layer is structured in your case, right? As XSDLC is the CRD that you have implemented.

So, maybe this is better stressed here: "Our XSDLC composition ..."

**[344]** (Text) again overflow here

**[345]** (Text) what does this mean? Please clarify in footnote


## Page 77

**[346]** (Text) Is this the actual implementation that you provide to realise all the magic?

If I understand well, then you just rely on the existence of the two providers to do your (reconciliation) work. 
Please clarify this. 

Question: does the CR also needs some kind of implementation (e.g., in terms of its content/schema). Or everything is covered in the Composition Pipeline?

**[347]** (Text) here the arrows showcase dependencies and order of creation or sth else?

**[348]** (Text) earlier? When was that attempted and when made you change your mind?


## Page 78

**[349]** (Text) What is this? It is not clear to the reader.

**[350]** (Text) gitops vs GitOps

**[351]** (Text) what is the semantics of these branches? Please clarify

**[352]** (Text) again this is not clear - very specific terminology that might not be clear to the reader.

**[353]** (Text) image pull of what?


## Page 79

**[354]** (Text) ok but what is a check? Does it relate to the principle checks that we show in the previous chapter or is it sth different?

**[355]** (Text) there is also the test-only check available in the CLI that is a combination of function-only and load-only. Should we consider that this is simulated by supplying its component checks?

**[356]** (Text) Overflow issue here

**[357]** (Text) which subset?

**[358]** (Text) which configures the validation? So, provides values for respective configuration parameters?


## Page 80

**[359]** (Text) overflow + now it is clear what is rate. But this load-test maps to a single load. So, it looks more as stability testing rather than load testing

**[360]** (Text) ok but the ordering is required by the model or it is enforced by the implementation (so at the modelling side, it does not matter)?

**[361]** (Text) So there are applied by default? Can't the user modify this?
As this should be configurable by the user, right?

**[362]** (Text) why aren't functional-tests also executed for the production gate?
Should we consider that as we have passed the staging gate already, then there is no need to perform them?
But validate-only is repeated in both gates! So, why not repeat and the functional tests then!


## Page 81

**[363]** (Text) so these two are fixed while the rest of the information is configurable by the user, right?
This could be clarified below.

**[364]** (Text) ok - but we should also expect that configurations can be given also by the CRD?
Can you provide an example configuration in this case?


## Page 82

**[365]** (Text) Are we missing per-environment resources here?


## Page 83

**[366]** (Text) , respectively,

**[367]** (Text) which helm chart? Is there Helm chart in your repo? Please clarify.

**[368]** (Text) is there any assumption about this installation? What should be already in place? A Kubernates cluster and Crossplane?

Please supply respective links with instructions to enable the interested reader to satisfy the assumptions before exploring your solution.

**[369]** (Text) please also provide reference to the figure that shows the CRD model.

**[370]** (Text) please validate whether it is 44 or 47. I have the impression that it can be 47 based on the rest of the sentence (the naming of all resources and the fact that some are replicated in 3 environments)

**[371]** (Text) overflow here

**[372]** (Text) ok - so it does the monitoring. The reconciliation is done by which function? The same?


## Page 84

**[373]** (Text) the arrow here has the right direction?
As the function should observe the resources and not the other way around

**[374]** (Text) by mistake or on purpose?

**[375]** (Text) can this be changed by the user? Or these are hardcoded?


## Page 85

**[376]** (Text) a


## Page 87

**[377]** (Text) (see Section 5.6.2 below)

**[378]** (Text) if there is a change, does it inform the previous component?
Shouldn't that component take action to reconcile the change?

**[379]** (Text) function-go?

**[380]** (Text) overflow + maybe this is moved to the description of the function?


## Page 88

**[381]** (Text) please be aware of the overflows which happen in almost every page


## Page 89

**[382]** (Text) so, this is a kind of configuration point?

**[383]** (Text) Maybe in Chapter 2 could provide more explanations about them in the context of the non-expert reader.

**[384]** (Text) so, these are the global configurations/settings ...


## Page 90

**[385]** (Text) default or the one that has been applied during the installation (as defaults can be overridden).

**[386]** (Text) this was not mentioned before.
So, it is also not clear what was involved in that tier and what was the respective philosophy.
So, maybe this sentence could be deleted?


## Page 91

**[387]** (Text) could this be an extension of your work to satisfy different potential "business" users?


## Page 93

**[388]** (Text) mapping to

**[389]** (Text) this is the name of the step?
I am asking as the name is given in an example that follows and is different.

The same for the other check types.

**[390]** (Text) overflow again here

**[391]** (Text) (PersistentVolumeClaim)

**[392]** (Text) this runs where?


## Page 94

**[393]** (Text) against the live API?

**[394]** (Text) chart installation?


## Page 95

**[395]** (Text) Maybe clarify that here you start indicating how the workflow template is dynamically generated (in terms of its core content).

**[396]** (Text) For example, ...


## Page 96

**[397]** (Text) so these will be some checks depending on the validation mode (default or configured).

**[398]** (Text) ok but the thresholds apply mainly for the load-test check. So, this holds only for that. For the other checks, we do not have thresholds but other kinds of config. params, right?

Final: did not say sth about functional-test check. What does it hold for this one?

**[399]** (Text) did not explain how this was produced and what is its content.


## Page 97

**[400]** (Text) and applied?

**[401]** (Text) I presume that the installation covers the first three steps. 
The application of custom resource is a responsibility of the user, right?

**[402]** (Text) Yes but this is not installation step per se and should not be advertised as such. 

This is a post-installation step applied by a user that attempts to exploit your solution.


## Page 98

**[403]** (Text) Nothing is mentioned about Sections 5.2.


## Page 99

**[404]** (Text) if each element maps to a section, could reference that section in parenthesis.

In overall, the structure of each chapter must be provided within the auspices of an introductory paragraph.

**[405]** (Text) in Chapter X ...

Or is this done in this chapter?


## Page 100

**[406]** (Text) The aforementioned

**[407]** (Text) which is guaranteed? Please clarify

**[408]** (Text) Maybe the vertical labels should be moved to the middle of each rectangle.
This will make it clearer where they belong.

**[409]** (Text) A has been "consumed" in the figure. Please fix this (Argo Events)

**[410]** (Text) The arrow here is correct?
Same for arrow connecting Sensor and Argo Workflow.

**[411]** (Text) The rectangle length should be reduced so as to be within the overall boundary. Now it looks like crossing boundaries, which is a "wrong signal"


## Page 101

**[412]** (Text) As the ...

**[413]** (Text) and by knowing the aforementioned URL pattern

**[414]** (Text) overflow ...


## Page 102

**[415]** (Text) If you refer to Section 7.3 that more or less conducts the latency analysis, the current section could be removed.

But I leave it up to you if you desire sth like that. In overall, it could be also argued that it is not bad to provide a short summary of the insights that Section 7.3 supplies so as to impress the reader with the outcome (latency is just fine with the current complexity that the overall environment has).

**[416]** (Text) ok but in principle, the test time can be greater depending as you indicate on specification size and validation mode but also on test coverage and intensity plus test types (e.g., performance testing could include multiple kinds like load testing, stress testing, stability testing) that could significantly increase the test/validation time. 

Thus, what is really important is the duration until workflow submission.
The workflow execution duration is also interesting for other purposes (how quickly can we validate the OpenAPI specification, how quickly can we validate that the code update is correct).


## Page 103

**[417]** (Text) Is there a reference that also clearly indicates this - that it is a bad practice to use a single repo?

It is optional to provide it as your arguments in the paragraph are quite strong.

**[418]** (Text) Lots of overflows in the current page


## Page 104

**[419]** (Text) As each ...


## Page 105

**[420]** (Text) This is a very big advantage.
I would "advertise" it as such in order to highlight this for your system!

I propose to collect all these strong advantages and do the advertising at the beginning, when your system is introduced and at the end (last chapter) when you provide conclusions for your system.

**[421]** (Text) I would recommend to better explain how hydration works in Chapter 2 (what is covered by the developer, how hydration works, how it can be changed, which CI/CD scenarios can be covered, etc.). 
This is a very important information for the reader in order to understand exactly what you are offering ..


## Page 106

**[422]** (Text) Figure is nice.
Two improvements:
(a) "manifests ..." label could be placed centrally in the central arrow that has to become more lengthy
(b) in overall some arrows lack their straight line and create a small bad impression about the figure.


## Page 107

**[423]** (Text) the labels here overlap!

**[424]** (Text) The arrows in this rectangle and the "DDT ..." one are misleading. Logically speaking, their direction is wrong. It should be the other way around (DDT enables to move from Development to Staging and not the other way around -> the same from Staging to Production)

**[425]** (Text) again overflow here. Please note that due to the overflow in many pages the information is cut and not completely shown.
Please correct.

**[426]** (Text) must be formatted?
What does it mean that it must match these two things?


## Page 108

**[427]** (Text) ok but you need to stress that the order of the environments is important as it is assumed a specific transition order between them (1st->2nd, 2nd->3rd, etc.).

**[428]** (Text) which workflow?

**[429]** (Text) again this is nice.
The vision could be to support other providers apart from GitHub ...


## Page 109

**[430]** (Text) this could be also shown with a nice figure, if possible

**[431]** (Text) very nice section - this is what I would like to see as a reader and as a potential adopter of your solution


## Page 110

**[432]** (Text) which signifies automated sync.

**[433]** (Text) ; it signifies ...

**[434]** (Text) ; this signifies ... along with the existence of a gate..

**[435]** (Text) but this can be also applied. There is no quality gate specified and autoMerge: false.

**[436]** (Text) ok but does it make sense? Which cases could it cover? Maybe it is not recommended as a potential configuration of an environment?


## Page 111

**[437]** (Text) Again some arrows do not have straight lines

**[438]** (Text) ok but one configuration (no gate, no auto-merge) is not covered by the figure. The question is whether it makes sense to cover it.


## Page 112

**[439]** (Text) functional testing is not covered in the sentence

**[440]** (Text) usually the Setpoint is compared against a variable. What is the variable in your case?

**[441]** (Text) I do not have a problem with OpenAPI specification being the Setpoint. 
However, there are two issues to consider:
(a) is the specification reflecting the current implementation or it refers to previous versions of it? We need to rely to an up-to-date OpenAPI specification. I am not sure this is guaranteed
(b) is the OpenAPI specification complete in terms of information regarding testing the sw? Again, this is not guaranteed and you have specific checks to check/assess it.

Thus, both of the above issues jeopardise and create risk in terms of your decision to consider OpenAPI specification as the setpoint.

**[442]** (Text) (approve) -> he/she should manually approve the promotions. This is a major developer feedback, right?

**[443]** (Text) The feedback in control theory signifies if the current situation is ok or not. So, it compares the process variable with the setpoint. Thus, maybe the feedback is the test report? Maybe with the addition of the developer that supplies manual feedback to cover uncertainties (should we have or not the promotion).


## Page 113

**[444]** (Text) Nice figure but there are issues as labels are not shown well and arrows are hidden.
Please improve to make it more professional

**[445]** (Text) correct - but based on my previous comments, the specification must be complete and up-to-date in the first place.

**[446]** (Text) ok but the issue is that the workflow has failed. Shouldn't it be re-run in order to get the feedback?
If we have a failure status, this gives a wrong signal. The developer will get it and will think that sth is wrong with his/her code.
I am not sure this is a desired situation.
The question is whether there can be any mechanism to correct/fix this.
Someone checking workflow executions and re-running workflows if they have failed.

**[447]** (Text) these are not the three well-known control-theory properties: stability, controllability and observability. 

It seems that you are mixing formal system engineering principles with control-theory.

**[448]** (Text) ok but the question is whether this holds for the setpoint.
The checks that you incorporate indicate the opposite!


## Page 114

**[449]** (Text) ok but if document(ation) is the pointset, it has to be complete.

But I agree that based on the way the DDT-adjacent principle has been defined, you need to verify the documentation.

But in control theory and based on the completeness property that you have mentioned, the documentation should have been complete.

**[450]** (Text) not clear what you are trying to identify here. Sth mentioned in the documentation that maps to a file not present in the source tree?
And how do you know if the documentation talks about a source code file?

Of course, this discussion depends on what you mean by documentation.


## Page 115

**[451]** (Text) main implementation logic (as now we talk about source code and its CI/CD)


## Page 116

**[452]** (Text) what is KubeCore?
I am not sure this was indicated before.
Maybe dedicate a footnote to provide the necessary explanations.


## Page 117

**[453]** (Text) (so only all relevant resources)


## Page 119

**[454]** (Text) Did not say sth about Section 6.8.


## Page 120

**[455]** (Text) If PoC evaluation does not lead to assessing various metrics, I would recommend following a different presentation order for the three evaluations conducted: first present the PoC, then the controlled experiment and finally the large-scale evaluation. This is more natural as we move from PoC towards real-world evaluations.

**[456]** (Text) Ok but aren't the evaluations also showing this?
That an ontological specification lead to the deployment and execution of quality-gated delivery pipelines?


## Page 121

**[457]** (Text) Question: do we see the PoC operational evaluation? Is this this last arm?
I am asking this as the title of the arm signifies sth different.

**[458]** (Text) detection accuracy is measured by which metrics? Before presenting the controlled experiment and its results, you need to define the respective metrics and explain how their values were computed.

**[459]** (Text) complete -> in the sense of specification completeness, which is necessary to guarantee determinism and observability based on your proposition.

Please clarify this here.

**[460]** (Text) what is meant by a task?

**[461]** (Text) are you sure about that?
I believe that the large-scale evaluation does not rely on this API!


## Page 122

**[462]** (Text) could explain how this is done.Do you have a different branch per defect injection? Or you have specific code that deliberately conducts the injections?
I presume the second if the injections target the specification-based principles (not the testing ones).


## Page 123

**[463]** (Text) is this expected?
I mean that you indicate here what is the expectation by evaluating your solution. 
Is this true?
If yes, then please modify your wording to reflect this. 
If not, then please explain what you did: e.g., I have used the CLI to conduct all principle-based checks and validation accuracy was perfect as each injection was the one solely discovered.

Need to also explain how the actual assessment was done: did you check the validation reports to see that indeed only the actual injection was detected and nothing else?

Final thing: maybe this minimal evaluation could become even stronger if you could also check all combinations (or even some of them and also all) to see that the detection accuracy is again correct?

Another final thing: if the injections targeted specific tests/checks per principle, then the validation might have been considered incomplete. In this sense, it could be extended to cover well/completely each principle. 
That is introduce issues per principle injection that target each a different check. And then see that indeed all checks of the principle identified the respective issue.
In my view, this extension is necessary in order to showcase that each principle is covered completely by the tests/checks and that the validation report produced is surely complete.

**[464]** (Text) ok but what is really meant by quality dimension?
As we know that there are nice principles. So, it is expected that all 9 principles are somehow covered in this experiment. But you have 5 APIs and not 9 ...

**[465]** (Text) Very interesting aspect to include OpenAPI version in the table.
However, you seem to focus mainly on 3.X versions. Thus, it would be nice to attempt to also cover 2.X versions to showcase that your work is fully operational also for them.

**[466]** (Text) If you have removed a specific quality dimension from here, the naming of the API should be consistent with this removal. 
If you provide the name perfect-api, this means that the API is perfect and does not exhibit any issue.
Based on what I see in the next table (7.4), it seems that the API does not exhibit any critical issue, so only 3 principles were affected (P002-P004). Thus, I propose to rename it as: non-critical-api in the sense that it does not exhibit any critical issue.

**[467]** (Text) Here the principle affected seems to be P005. The same holds for the next API. But you said that each API differs from the other in terms of affected quality dimension. So, I do not really see here what is the real difference between bad-docs-api and no-auth-api. I understand that the target defect class is slightly different but the principle affected is not. 

My suggestion: have one from these two APIs focusing on P001. 
In this sense, the principle affected will be different but still both principles will be critical.

**[468]** (Text) ok but it is not clear how degradation is detected in this case. You put a specific on-purpose delay. However, there should be a threshold that needs to express the respective latency requirement.
So, please indicate how this threshold was communicated to the system. Maybe in the CR?

**[469]** (Text) ok - are these two also critical or not?
Because you state in the text below that there are blocking and non-blocking failures. Blocking means critical, right?

**[470]** (Text) but in the table, maybe all defects seem to be blocking?


## Page 124

**[471]** (Text) I disagree here with the table content in the sense that it has to cover everything, I mean that all principles need to be covered as this is the ground-truth for the controlled experiment in the sense that it clarifies whether indeed the quality gate system is working properly. 

Further, the table could be utilised as the ground truth for the validation as it can really indicate whether all injections in all APIs were indeed detected by the CLI.

**[472]** (Text) I am puzzled!
The perfect-api as was presented in previous section does not exhibit any principle-specific issue. 
Yet here it seems to have issues in three principles.

**[473]** (Text) In the previous table, you indicated that bad-docs-api suffers from a security-oriented fault. However, here, it seems that the principle is passed. This is a major inconsistency in my opinion. 
Please correct it.

**[474]** (Text) but as I understand, this is not the perfect-api that was mentioned in the previous section. It is another version of it that has issues in P002-P004. 
So, its name must be modified.

**[475]** (Text) By reading this paragraph, I am more puzzled than before about what is the perfect-api. 
Previous section has shown that perfect-api is well-designed and does not have any principle-specific issue. In this section, you indicate that even "perfect-api" is derived from the perfect-api of the previous section by injecting some faults on it. 
Then, in this paragraph, you seem to indicate that perfect-api, i.e., the original API, while well-designed, has quality gaps, which does not make sense!

**[476]** (Text) I would indicate that non-critical-api has been simulated to match current development practice that does not cover all quality dimensions in API specification.

**[477]** (Text) no-auth-api was said to be missing authentication scheme and bad-docs-api a security scheme.

**[478]** (Text) so these were the two faults injected in the broken-api?
I believe that these should have been communicated earlier in the text of this section.
The same could be done for all APIs (other 4) that were derived from the ideal one.


## Page 125

**[479]** (Text) yes but this must have been already communicated.


## Page 126

**[480]** (Text) I propose to have 4 layers so that you have another evidence that the gate mechanism operates well even in the case that the OpenAPI is not valid. So, one of the bad-docs-api or no-auth-api could be made to have this problem in order to showcase that test-ready mode also works and is handy.

All modes should be handy and well-functional but also serving distinct purposes & requirements.

**[481]** (Text) 4/9 as there are 4 principles that are critical

**[482]** (Text) I believe that as you mention slow-api that concerns the performance testing, similarly the broken-api should be also mentioned as more or less exhibits the same pattern: does not fail the staging gate as it exhibits warning issues but fails afterwards as it cannot pass the functional tests (P006).

**[483]** (Text) ok but still the issues that exhibits can influence functional testing. Thus, someone can argue that many more issues/principles are critical.
In my view, it can be the case that a principle carries both critical and non-critical aspects. So, the evaluation of the principle can lead to deciding whether the current failure is critical or warning. 
So, it can be the check, part of a principle, and not the principle itself that impacts the final result.
Unless the principles are organised in such way that each principle carries only one type of issues, either critical or warning.


## Page 127

**[484]** (Text) In this respect, the idea is that there are two cases that can enable promotion from deployment to staging:
(a) no critical issues & warning issues occurr -> automatic transition
(b) no critical but warning issues occur -> transition only upon approval.
This can cater for more flexibility but necessitates a lack of sth, potentially another configuration parameter or the use of conditional attributes. I would propose the second alternative: instead of saying just merge automatic or manual, you can have a "grey" value and indicate the following: there should be a condition that distinguishes between automatic & manual (based on what I proposed earlier). Such a condition could be expressed by considering the existence or not of issues. For instance, in case of staging-to-prod, the condition could signify that if there are no issues at all, transit directly; otherwise, let the developer decide. In case of dev-to-staging, if there are warning issues in P00X and P00y, let the developer decide. Otherwise, If there are warning issues in P00z and P00T, transit automatically.

Thus, the developer can specify conditions that concern both the issue criticality and the source (principle). 
To me, this is more powerful and expressive than the use of a single switch (automatic or manual).

**[485]** (Text) Need to provide precise number of APIs downloaded.


## Page 128

**[486]** (Text) The OAUth2 scope is needed when type: oauth2 (or OpenID Connect built on top of OAuth2).
Is that the case with the current OpenAPI specification?
I ask this as you mention in general that OAuth2 scopes are needed but the question is whether this is relevant in the current case.

**[487]** (Text) I do not like this title here for the column. As you have done the evaluation, it pass rate should be exact and not approximate. 
Unless it is an estimate of an LLM that can be very imprecise depending on the context!

**[488]** (Text) What does high, low and medium mean in practice? High is >= 90%, medium is >= 50% and low is <50%?

Need also to supply raw numbers, so how many APIs passed P001, P002 ...

**[489]** (Text) It would be nice to move the focus also on each principle in order to see which checks fail under what percentages. This would be a very interesting direction that could unveil even more interesting results

**[490]** (Text) Did not give any comment about P005 and where the pass rate could be considered as satisfactory. 
Here the analysis could be even more interesting if we could discern between essential lack of information and test-oriented (information that can impact testability). In case of the former, things are really bad, highlighting that many APIs have significant quality issues that can even impact their actual selection and usage. 

Please provide a respective paragraph in the text.


## Page 129

**[491]** (Text) partially in my opinion: it matches P001 but matching fails for P002-P004 in the sense that 0 is different from low. 0 really means nothing. Similarly, medium is different than 80% that should be high. 
In any case, there is no strong need to match one evaluation over the other as each has a distinct purpose. 

If you believe that this correlation is important to be shown in order to strengthen the previous experiment, then that experiment will need to be extended so as to simulate more or less the real situation.

**[492]** (Text) were not passed by any of the APIs

**[493]** (Text) to me 80% is high rather than intermediate.

**[494]** (Text) I believe that there is a need to not provide similar information as already supplied in previous chapters but to focus on fully describe this PoC deployment.

Currently, we have very little information about the experiment configuration (what was in place) and goals while it is also not clarified what was done during the experiment. We only see in last sub-section some results which are justified in one column with no further analysis plus references to some repos/projects that were never introduced. 

Please re-organise this section accordingly.

**[495]** (Text) Please explain the structure of this section in one paragraph and correlate it with the section goals.

**[496]** (Text) which one?


## Page 130

**[497]** (Text) ok but of course the transition from one environment to the other requires manual feedback in some cases. So, there is high automation degree in all parts where this is necessary.

**[498]** (Text) What do you mean by "declared specification"? I believe that what you are trying to do is to evaluate the completeness of the OpenAPI specification. Such that the specification can be used for the real automated testing of the respective API.

**[499]** (Text) This relates to P006 and P007 while here you talk about the other principles that relate to the quality/completeness of the OpenAPI specification!


## Page 131

**[500]** (Text) API testing isn't a form of integration testing?
So, here you mean "other kinds of integration tests"?

In addition, unit testing is of course a developer's responsibility. However, integration and other kinds of more composite testing go to the CI/CD realm. Thus, in principle, they must be part of a pipeline that should be executed by your solution!

**[501]** (Text) Indeed, but in this kind of testing, not just functional (API) testing but many sorts of testing are applied, including unit, integration, contract, end-to-end, performance, load/stress, reliability testing plus various sorts of security testing nowadays.

Thus, DDT should be extended to cover them all!

In addition, apart from performance, security thresholds should be given while also coverage thresholds might also be considered.
These are all criteria that can influence the final decision, if we need it to be automatic.

**[502]** (Text) sure but instead of load testing, maybe it is better to do stress testing?

**[503]** (Text) Ok but may this leads to the requirement to have two control loops instead of one as you make the system even more dynamic by changing the setpoint.


## Page 132

**[504]** (Text) Not clear how these were derived.
Need to explain the respective rationale.
Maybe also a reference would further strengthen their consideration.

**[505]** (Text) Ok but what does it mean pass and fail? 
Pass means that the problem passed the control but it shouldn't?
Or does it mean that the issue was addressed successfully?

Please clarify in the text.

**[506]** (Text) Is this really an issue? As the system has been designed deliberately to address/cover it.

**[507]** (Text) I presume that this scenario shows that the system can scale to cover multiple repos / CRs ...

**[508]** (Text) you mean the ones used in the first experiment/validation? Please clarify

**[509]** (Text) I assume that this covers a complete history / period of real usage of the repos. So, instead of having just one commit per repo, multiple were issued. 
Maybe also indicate approximately how many commits were applied per repo?

In addition, the real question is whether the system exhibited the right behaviour. This can mean various things  like is the phase distribution correct?

**[510]** (Text) These projects/repo were never introduced before.

I believe that there is a need to fully describe the initial situation and the way the experiment was made before presenting any result!!!

**[511]** (Text) So, this is sth that needs to be corrected in the future? How this could be done? Is it possible that there are conflicting actions that could be done by different components, which cannot be avoided?
The main issue that the system always reaches a stable, valid state.
This should be always guaranteed.


## Page 133

**[512]** (Text) Please also indicate how these metrics were computed. If multiple tools were utilised, please mention them.

**[513]** (Text) you mean of the tool when run in standalone manner?

Maybe you can add a column that explains the metric semantics or can add further details in the notes when needed (only for some metrics some things are not clear like the current one)

**[514]** (Text) This could be inspected by making observations during the large-scale experiment. There the specs have varying sizes so you could separate the specs into low, medium and big-size ones and then provide validation time ranges per each spec class.

**[515]** (Text) does this depend also on the computing power of the cluster? Or can we consider it as more or less stable?

**[516]** (Text) Could indicate here the factors that can influence this metric as done for validation time

**[517]** (Text) could you have respective estimates for manual configuration? This could make a very interesting comparison!!!

**[518]** (Text) (Related to RQ1)

Please put this in each relevant parenthesis as otherwise you make the impression that you just directly refer to the main evaluation goals (e.g., RQ1 is called methodology soundness).

**[519]** (Text) Seems that paragraph is not fully justified

+

Of course, the experiment could be extended per check/checker basis to cover all possible cases per principle.


## Page 134

**[520]** (Text) Correct API addressing & classification?

I propose just a new title to strengthen the result in concert to the content of the corresponding paragraph.

Correct API addressing surely relates to the soundness of your methodology. 
While discrimination as a word does not convey the right things ...

**[521]** (Text) Correlation to RQ2 is not quite evident and justified in the paragraph. Please try to strengthen this. Of course, this correlates to the automatic transitioning between stages ...

**[522]** (Text) it depends on what is inside the layer. If the layer covers everything, then it has to cover all defects. If it covers a specific aspect, only particular defects apply.

**[523]** (Text) which are the 4 defect types? Invalid API (P001), critical quality issues (P005), functional faults (P006) and non-functional faults (P007)?

**[524]** (Text) Maybe change to: "Automation extent confirmed due to validation approach generality" or "Validation approach generality".

Further, in essence, this result relates to both RQ1 and RQ2: The methodology is sound and can be generically applied & this generalisation enables to achieve better automation ...

**[525]** (Text) I do not understand how quality gaps connect to RQ2 which is about automation extent. Automation extent is about how extensive is the automation achieved by your solution. This has nothing to do with the quality gap, which can be detected only based on RQ1 and the way it is implemented.

**[526]** (Text) ok, what does specification mean?

**[527]** (Text) RQ3 + RQ4 as RQ4 is also proven/validated based on PoC as the ontological specification enables to produce and maintain a running system which applies the proposed methodology


## Page 135

**[528]** (Text) Please update the chapter to better clarify the own effort that you have put in the thesis.
Clearly indicate: percentage of code that you have written, how much time was spent in writing and updating .md files, how much time you spent checking AI-authored code, documentation & thesis, how many times did you intervene in the process to correct things and provide feedback, in what forms was this feedback supplied.
Please also indicate whether did you make internal manual checks over each final or almost-final artifact produced (code, documentation, thesis) and across them apart from the checks that were made automatically by the agents.
Finally, please indicate the following: (a) what is the percentage of authorship that you can claim in overall over the whole thesis/project, (b) whether you believe that you have fulfilled your role well and guided correctly the agent system - which role could be regarded to be undertaken if this work was done in the context of the sw project? Would you be an analyst, an architect?
Very importantly: your original idea was fully implemented in the end or there were major deviations? Were these deviations due to certain issues or because you found that the pathways suggested or followed later on where more interesting?
(c) what else you could improve in the near future in following projects, what is the own experience that you have gained? Not just the lessons learned by the agents themselves.

**[529]** (Text) software


## Page 136

**[530]** (Text) ?

**[531]** (Text) why do you say "later" here?
What was the initial situation?


## Page 137

**[532]** (Text) Ok but maybe you could explain why did you follow this structure and how many levels were needed to achieve the final outcome?
Are there any guidelines that indicate how someone should structure software projects and accompanying artifacts like documentation ones?
Or this was done purely based on your current experience?
Please clarify in the text.

**[533]** (Text) ok but each file implements a specific principle, right? Please clarify.

**[534]** (Text) I do not understand this as previously you have indicated that an agent acting on a leaf .md has access to the .mds of its ancestor dirs.
So, why to repeat information/knowledge that already exists and will be essentially available to the agent?


## Page 138

**[535]** (Text) I would indicate below in text which are these code quality rules.

**[536]** (Text) where do you indicate which principles are critical or not? How and where validation modes are related to the principles?

**[537]** (Text) so, by API inventory you mean the repo with APIs that you have collected from API Guru? While evaluation purpose indicates which kinds of evaluations had to be done on the repo and which results would need to be reported?

**[538]** (Text) this includes figures you would like to have apart from those to be generated by the agents?

**[539]** (Text) Not clear what is to be included in the documentation. Please clarify. The diagrams are UML models or sth else?

**[540]** (Text) CSV is used for what?

**[541]** (Text) With results you mean evaluation results or sth different?


## Page 139

**[542]** (Text) so the thesis is considered as a kind of documentation or even specification? Otherwise, it does not make sense to update code if some analysis is wrong. It also depends on what is this analysis and how it connects to the code. 
Can you provide an example here to better justify your choice? As well as an even more convincing argument why thesis can be considered as specification?

**[543]** (Text) ok - I understand the three first principles.
But the last one was really applicable? Which directories were produced by the AI agents themselves (rather than predicted by you)?

Please note that I agree with the rule.

**[544]** (Text) to which category of memory systems it belongs? Please clarify.


## Page 140

**[545]** (Text) ok, so this information is complementary.
I presume that this is needed for other reasons rather than surviving from crashes, right?
E.g., it can be used to record lessons learned such that they can be utilised also in other projects.

I see some answering below but maybe it is better to clarify this from the very beginning

**[546]** (Text) (between the agents)

I assume that is the current case here ...


## Page 141

**[547]** (Text) ok but these are precise rates. As in the evaluation chapter, it was mentioned sth like estimated rates, which is sth different ...

**[548]** (Text) what was wrong? Please provide your experience here ...

**[549]** (Text) Of course, I would expect that the thesis was written by you. And that the agents could check it against the code base to find inconsistencies and report them to you.
Alternatively, parts of it could be written by the agents and you could then read them to make the necessary corrections. This feedback could then be given back to the agents to correct the code ...

Of course, in real projects, more automation is needed as time is money so if documentation is produced in this way, this is of course more than welcome. Provided that the documentation is indeed correct and someone reads it.


## Page 142

**[550]** (Text) Ok, theoretically I see that there is a connection. It was also proven by the AI-based development that feedback leads to improved version of implementation & documentation.

However, what I am missing is a prove that indeed the feedback from DDT leads to corrections. This could be part of the evaluation: e.g., based on the feedback of one problematic API, you could supply this to an agent to see whether it can improve its specification. In this way, the theoretical connection is also proven empirically ...


## Page 143

**[551]** (Text) ok but this should have been evaluated more extensively to be sure about it ...
All checks/tests should have been considered and not just 1-2 per principle


## Page 144

**[552]** (Text) returned?

**[553]** (Text) but was this the case? Did you also make manual changes in code in later commits? Or only at the beginning?


## Page 145

**[554]** (Text) deadline pressures make artificial agents do mistakes?


## Page 147

**[555]** (Text) Please note that OpenAPI v2.X support was not examined in the evaluation. This is another threat to validity in my opinion. 

Please collect any relevant issue I have detected in previous chapters (especially the evaluation one) in order to present it here in this chapter.

For instance, the expression of quality gates could be more powerful. So, it is not clear whether it could cover all possible requirements and preferences.

The reliance on single kind of performance testing + limited functional testing is another major issue.

**[556]** (Text) threats to validity ...
Are these also covered?

What is the structure of the chapter?

**[557]** (Text) As I have already stated, please checks whether the term methodology that is too strong can be replaced with something lighter that fits better your contribution. 
Soundness will then require less to be proved wrt a formal methodology.

**[558]** (Text) the part "..."

**[559]** (Text) but functional tests don't they cover business logic errors?
The API exposes an interface that is not properly implemented in the code due to business logic errors.


## Page 148

**[560]** (Text) correct but it has to be mapped well, precisely and formally to these layers within the manuscript. And this needs to be done from the very beginning, i.e., the chapter dedicated to unveiling these 9 principles.

**[561]** (Text) I believe that you did not comment on the part about quality gates. 
In essence, the specification part is not just the OpenAPI description of the API but also the CR. 
By accounting also the CR, then it could be indicated that quality gates can be derived. 
However, across the whole report, the CR is not accounted for ... only to highlight the ability of the system to create a Kubernetes-based closed control-loop that enables the gate-based transitioning between environments through the automatic execution of CI/CD pipelines.

Thus, if it would be part of the specification from the very beginning, then you could claim that quality gates can be automatically derived.

**[562]** (Text) fully what? This research question was addressed fully?

**[563]** (Text) OpenAPI

**[564]** (Text) correct - but the business logic is already specified in the OpenAPI specification both structurally and descriptively via textual descriptions. 
The scope is also to functionally verify the API, not a whole application.
The issue is how rich is the examples parts in the OpenAPI specification and whether you could adopt already established API testing approaches that cover more advanced API testing scenarios.

In other words, you touch the subject but not deep enough. So, future work has to improve upon it.


## Page 149

**[565]** (Text) Could refer to the specific section with the large scale evaluation to be even more precise here.

**[566]** (Text) could be removed + when this would take place?

**[567]** (Text) Totally agree here. But I am not sure this was mentioned in the evaluation.
So, this is your own assumption based on your experience and your knowledge about how your system works and integrates with external ones. 
Please reflect on this thinking in the current paragraph.


## Page 150

**[568]** (Text) but the finding seems to be negative while here you describe something positive ...
So, this is not reinforcement per se. Please correct.

**[569]** (Text) logically speaking as part of a complex control-loop architecture with two levels ...


## Page 151

**[570]** (Text) did you forget security here?
Otherwise, the math afterwards is imprecise as we expect 1/5 and not 1/6. So, one principle is missing ...

**[571]** (Text) its

**[572]** (Text) the

**[573]** (Text) I agree with you here.
Maybe you could add here that maybe the existing tooling allows to create correctly-structural APIs but the time pressure does not allow to enrich them (e.g., via manual annotations in code or manual enhancements directly at the produced OpenAPI specification).

**[574]** (Text) at least the most critical ones (e.g., from staging to production).

As you also shown in the evaluation, warning issues do not impact the promotion from development to staging.
Unless this is enforced by the organisation itself or the project manager. 

Please do not also forget the security aspect, which is critical.
Neglecting it surely will deter any kind of promotion!
This is the core difference between critical and warning issue detection. The former is blocking while the latter could be semi-blocking ...

**[575]** (Text) Please supply a respective reference here (to strengthen your argument/claim).


## Page 152

**[576]** (Text) Totally agree here.
In fact, the developer could manually do the integration or include it as a later step in a CI pipeline.

So, it is his/her own responsibility to do that and attempt to automate it in the correct way ...

Another thing that came to my mind: what is a good practice in the context of branching that should lead to optimally exploiting your system?

For instance, if someone want to reach until staging/testing for some branches and production for others, how could optimally do that? Would he/she require to create one or multiple GitOps repos?
These are the kinds of scenarios that someone would like to know before deciding to exploit your system.

**[577]** (Text) However, this ...

Ok but this is a valid scope restriction, especially in the context of B.Eng. thesis!!!

**[578]** (Text) Correct! This was my own observation.
In addition, did not evaluate whether all possible issues are detected within a single principle.
We have two levels here: multi-issues per principle and multi-principle issues.
Please indicate this by extending the current paragraph.

**[579]** (Text) I agree here.
The issue with the integration between an external Git service and your system is part of external validity? Or internal? Please classify and include it accordingly.

**[580]** (Text) or risk-based scoring model? It matches better the first option of severity-weighted. However, risk is sth more complex but also more valuable ..

**[581]** (Text) sure but this is your own responsibility here.
So, the question is whether it was undertaken and in which degree!
Plus whether there was also a kind of subjectivity also from your side and for what reason.
Of course, this subjectivity should have been avoided ...


## Page 153

**[582]** (Text) correct - but still your system can also detect functional and non-functional issues, even limited. So, this is also worth mentioning ...

**[583]** (Text) ok but please decide. The benefit is for the platform team, the developer team or both?

Once you decide, please also correct the title, if necessary.

**[584]** (Text) ok - I do not disagree.
But what is "meant" by model here? This has to be clarified for the sake of the reader. 
Is this the hierarchical model of the project with the CLAUDE.md files being placed in every suitable directory?


## Page 155

**[585]** (Text) again, please check the term as it is too strong for your contribution.

**[586]** (Text) validity-mode-aware ...

**[587]** (Text) This is a ... which was also developed in the context of this thesis.

**[588]** (Text) Based on this CRD, a ...


## Page 156

**[589]** (Text) This is a ... or This represents a ...

**[590]** (Text) ok but your answer does not cover the "comparable" part. Is it really comparable and why?

**[591]** (Text) Again, "fully" is not sufficient - you need to explain what is "fully"!

**[592]** (Text) standard

**[593]** (Text) This was not so well advertised. In addition, I am not sure that it includes all appropriate checks that cover automated testability ...

**[594]** (Text) Again, did not cover the "quality gate" part in the RQ ...


## Page 157

**[595]** (Text) The question is how effectively.
So, your answer here should provide a degree ...

**[596]** (Text) OpenAPI specification?

**[597]** (Text) correct - did the evaluation proved that and how?
For instance, does the reconciliation always work as expected?
Can it sufficiently address any deviation from the desired state? This is the main question to also answer here.


## Page 158

**[598]** (Text) quality-gated pipelines can ...


## Page 159

**[599]** (Text) ok but is it sufficient? This is another major question here. I guess not. So, this must be communicated.

**[600]** (Text) Agreed. But even if the versioning strategy is described in the OpenAPI specification, I am not sure that your current method can properly check it. It is too restrictive and too specific - it cannot cover all possible cases. This is another major limitation in my opinion.

**[601]** (Text) Plus the need to cover other kinds of non-functional testing, including security ones.

**[602]** (Text) the other evaluation with different API versions is also limited ...

**[603]** (Text) I would also add the expressivity of the quality gate conditions - currently,  it is more or less fixed. I have proposed a way to enrich it.

**[604]** (Text) I totally agree here. The question would be whether only this could be adjusted or even other things related to non-functional aspects (e.g., security or changing load percentage over different API versions in canary-based scenarios) or even functional (e.g., do more strict or more detailed testing with higher coverage as issues were identified or reported by clients).

**[605]** (Text) Please highlight this as the main vision: having a closed control-loop potentially with multiple levels that enables system self-adjustment!!! Potentially, operating across different clusters ...!!!

I would put this in a first paragraph and then provide the respective direction-specific paragraphs as pieces or steps towards the vision.


## Page 160

**[606]** (Text) Totally agree here. One additional thing: determining the right source of truth and updating the other counterpart accordingly. This is related to whether the OpenAPI spec is the "specification" or the implementation code (that e.g. modifies the API behaviour).
It is interesting to check this scenario as it can influence the correctness of the adaptation actions that can be taken. If your real source of truth is the code but you adapt it based on a less fresh OpenAPI specification, this is a wrong adaptation action!

**[607]** (Text) ok, correct.
But based on the promotion scenarios, adding these two principles in the strict mode will have the effect that it would not be possible any more to transit from development to staging environment!
Thus, the question is whether these two principles belong to this validation mode or need to move to another one even stricter.
Or maybe you need to allow more flexibility, allowing the user to select the principles to apply in any relevant mode. 
Think of it -- in any case, I agree, of course, that the report should include the findings from the P006 & P007 principles.

**[608]** (Text) correct but the devil is in the details. 
Would all principles be applicable to the different standards? Would all checks be applicable? Or is there a need to move the abstraction in each principle such that you can cover any kind of standard?
Please clarify here in the text.

**[609]** (Text) Very nice scenario - quite realistic ...!!!

**[610]** (Text) So, the goal here is to create higher-level operators? I am asking this as in principle, XSDLC can already act as an operator child resource.

Thus, the title of the direction is slightly misleading.

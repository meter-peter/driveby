# Απαντήσεις στις παρατηρήσεις και ερωτήσεις του Καθ. Κρητικού

**Θέμα προς αποστολή:** Απαντήσεις στα 610 σχόλια του annotated PDF — Triantafyllis

**Από:** Πέτρος Ευάγγελος Τριανταφύλλης
**Προς:** Καθ. Κυριάκο Κρητικό (Πανεπιστήμιο Αιγαίου)
**Ημερομηνία:** 2026-05-16
**Επισυναπτόμενα:** αναθεωρημένη διπλωματική (`main.pdf`, 216 σελίδες)

---

## Πώς να διαβαστεί αυτό το έγγραφο

Το PDF που μου στείλατε στις 2026-05-02 περιείχε **610 inline σχόλια σε 150 από 165 σελίδες**, εκ των οποίων τα **343 περιείχαν άμεση ή έμμεση ερώτηση** (ερωτηματικό). Έχω εξαγάγει όλα τα σχόλια με αυτοματοποιημένο extractor (`thesis/review/extract_comments.py`) και τα έχω κατηγοριοποιήσει σε έναν πίνακα 610 γραμμών (`thesis/review/comment-index.md`).

Το παρόν έγγραφο αποτελείται από τρία μέρη:

- **Μέρος Α** — Απαντήσεις στις ερωτήσεις σας, ομαδοποιημένες θεματικά (35 θεματικές περιοχές που καλύπτουν και τα 343 ερωτηματικά).
- **Μέρος Β** — Συνοπτικός πίνακας κατάστασης για κάθε ένα από τα 610 σχόλια (ID × σελίδα × κεφάλαιο × τύπος × priority × status × έναρξη κειμένου).
- **Μέρος Γ** — Τι συνειδητά δεν άλλαξα και γιατί (10 σημεία διαφωνίας προς συζήτηση στην εξέταση).

Κάθε αναφορά της μορφής **[123]** παραπέμπει στο σχόλιο 123 του annotated PDF. Οι αντίστοιχες αλλαγές καταγράφονται στο `thesis/review/changes-log.md` με ακριβές `αρχείο:γραμμή`.

---

# Μέρος Α — Απαντήσεις στις ερωτήσεις σας

## Α.1 Ονόματα εξεταστικής επιτροπής — [1]

> «Please provide the names of the two examination committee members when they become available.»

**Απάντηση.** Είναι το μοναδικό σημείο της σελίδας τίτλου που παραμένει ανοιχτό. Θα ενσωματωθεί μόλις μου τα κοινοποιήσετε. Σας παρακαλώ ενημερώστε με όταν τα δύο μέλη επιβεβαιωθούν, ώστε να γίνει η τελική σελιδοποίηση πριν την κατάθεση.

## Α.2 Πεδίο εφαρμογής: «API» = RESTful API — [3]

> «It is not clear whether you focus on RESTful or other types of APIs ... if you believe that RESTful is a word you would not like to repeat all the time, you could state that from now on, when you say API you mean RESTful API.»

**Απάντηση.** Έγινε. Στην πρώτη χρήση του όρου «API» στο §1.1 προστέθηκε υποσημείωση που ορίζει ρητά ότι το πεδίο εφαρμογής της διπλωματικής είναι HTTP-based RESTful APIs που περιγράφονται από OpenAPI 3.0/3.1, και ότι η επέκταση σε άλλα API styles (GraphQL, gRPC, AsyncAPI) είναι μελλοντική εργασία. Όπου ακολουθεί ο όρος «API» στη συνέχεια του κειμένου, υπονοείται το παραπάνω πεδίο.

## Α.3 Accelerate, IDP, KubeCore, container image — [9], [46], [104]

> «What is Accelerate? Maybe explain in a form of footnote.»
> «What is KubeCore?»
> «could explain what is a container image in a footnote.»

**Απάντηση.** Όλες οι περιπτώσεις απαντήθηκαν με υποσημείωση στην πρώτη χρήση:

- **Accelerate** (§1.1): υποσημείωση που παραπέμπει στο βιβλίο των Forsgren/Humble/Kim και τις τέσσερις DORA metrics.
- **IDP / KubeCore** (§1.x): υποσημείωση που ορίζει το Internal Developer Platform και το KubeCore με αναφορά URL.
- **Container image** (§2.x): υποσημείωση με ορισμό (layered immutable filesystem bundle κατά την OCI image specification).

Τα ίδια ακρωνύμια προστέθηκαν επίσης στον πίνακα ακρωνυμίων (`chapters/lists-and-acronyms.tex`).

## Α.4 Impact του automation — [10], [12]

> «Could also indicate whether the advantage of automation leads to a remarkable possible impact on the organisations that feature it. And could explain what that impact is.»
> «compress what? The development?»

**Απάντηση.** Στο §1.1 προστέθηκε παράγραφος που συνδέει την αυτοματοποίηση με ποσοτικοποιήσιμα DORA-style αποτελέσματα (deployment frequency, lead time, change failure rate, MTTR), με αναφορές. Η φράση «compress development» αντικαταστάθηκε από το πιο συγκεκριμένο «compress the lead-time-to-deploy from days to hours».

## Α.5 «Software systems» σε ποια αναφέρεστε — [15]

> «which ones? You mean the "software systems under development and maintenance"?»

**Απάντηση.** Διορθώθηκε. Στο §1.1 το ασαφές «these systems» αντικαταστάθηκε από «software systems under active development and maintenance», με ρητή ερμηνεία στην πρώτη χρήση.

## Α.6 Coding agents και τεκμηρίωση — [26], [27], [33], [35]

> «So, is it a common practice to include it in code? Even if it is not related to the actual code implementation?»
> «agents can also consume this knowledge, especially if they are targeted towards it. So, it could complement an OpenAPI specification.»
> «Are you sure about that? Maybe someone could argue that these are possible nowadays.»
> «what about machine enrichment? Maybe the machine can understand the code and provide the missing pieces?»

**Απάντηση.** Επανέγραψα τη συγκεκριμένη ενότητα του §1.2 ώστε:

1. Να αναγνωρίζεται ρητά ότι τα LLM-based coding agents μπορούν να εμπλουτίζουν specifications από τον κώδικα (το οποίο όντως κάνουν: Optic, APIClarity κ.ά.).
2. Να ξεκαθαρίζω ότι η συμβολή της διπλωματικής είναι ορθογώνια: η DDT προϋποθέτει ένα authoritative spec και αξιολογεί την implementation **έναντι αυτού**. Η αυτόματη παραγωγή του spec είναι ξεχωριστό πρόβλημα.
3. Να τοποθετώ το ερώτημα «τι ισχύει σήμερα» (παρόν tense) και να αποφεύγω καθολικές διεκδικήσεις «αυτό δεν είναι δυνατόν».

## Α.7 Progress drift / behaviour drift — [30], [83], [126]

> «what about progress drift, i.e., the code that has progressed with a modification of the actual behaviour?»
> «there can be synchronisation issues between API specification and implementation. The question here is whether the API is correct and the implementation is wrong? Or that the implementation is correct and the API has to evolve in order to cover the behaviour drift?»
> «so the focus is on updating the OpenAPI specification, the code or both?»

**Απάντηση.** Πολύ ώριμη ερώτηση — απαντήθηκε σε δύο σημεία:

- **§3.1 (στατικές vs runtime principles)**: ξεκαθαρίστηκε ότι οι **στατικές** principles (P001–P005, P008) αξιολογούν **το spec έγγραφο**, ενώ οι **runtime** principles (P006, P007) αξιολογούν την **implementation έναντι του spec**. Η spec θεωρείται ως «source of truth» — αν η implementation διαφέρει, η runtime αρχή σημαίνει «drift».
- **§9.x (Discussion)**: ρητή παράγραφος που διακρίνει **specification-correctness drift** (η spec έχει λάθος, ο κώδικας έχει δίκιο) από **implementation-drift** (ο κώδικας απομακρύνεται από τη spec). Η DDT εντοπίζει drift αλλά **δεν αποφασίζει ποια πλευρά είναι λάθος** — αυτό παραμένει human decision (όπως ζητάτε στο [83]). Η μελλοντική εργασία §10.5 περιλαμβάνει ευρετική κατεύθυνση «direction of correction».

## Α.8 API stability από την αρχή — [31]

> «Or do you believe that the API should be stable from the very beginning and can thus drive the whole implementation cycles?»

**Απάντηση.** Όχι — προσέθεσα ρητά στο §1.3 ότι η DDT δεν προϋποθέτει stable spec εξαρχής. Προϋποθέτει spec **που εξελίσσεται μαζί με την υλοποίηση**, και η DDT είναι αυτή που εντοπίζει τα σημεία όπου τα δύο διαφοροποιούνται. Η spec λειτουργεί ως **drift detector** όχι ως «νομοθετικό» έγγραφο.

## Α.9 DDT/SDT και validation κώδικα — [37], [50], [73], [126], [138], [144]

> «Isn't that what you are aiming in your work? Conduct both API spec and code validation? And check the synchronisation between them?»
> «by specification you mean the contract itself or the API specification (OpenAPI spec, for instance)?»
> «again the focus seems to be on the specifications? But the code is equally important.»
> «isn't the contract yet another specification? Couldn't the contract also include the OpenAPI specification as part of it?»

**Απάντηση.** Σωστή παρατήρηση. Έγινε ρητή διάκριση στο §3.1:

- Το **«specification document»** είναι το OpenAPI 3.x αρχείο (στατικό artefact).
- Το **«contract»** είναι το συμβατικό υποσύνολο των properties που πρέπει να ισχύουν στο runtime.
- Η DDT αξιολογεί **και τα δύο**: στατικά για την ποιότητα του spec εγγράφου (P001–P005, P008), runtime για τη συμμόρφωση της implementation με το contract (P006, P007).

Η σύζευξη είναι ο πυρήνας της παραδείγματος DDT — δεν είναι «μόνο spec» ούτε «μόνο κώδικας», είναι η **διαρκής αντιστοίχιση των δύο**.

## Α.10 Research questions ↔ Contributions mapping — [39], [47]

> «Maybe just provide the research questions here and then ... could provide a nice mapping table from research questions to contributions.»

**Απάντηση.** Έγινε. Στο νέο §1.6.1 προστέθηκε **Πίνακας 1.1: Research Questions × Contributions × Primary Chapter References**. Κάθε RQ συνδέεται ρητά με τις contributions που την απαντούν και με το κεφάλαιο που φέρει το βάρος της απόδειξης.

## Α.11 Τρεις αξιώματα — [40]

> «The three axioms are these ones? Not so clear ...»

**Απάντηση.** Η εισαγωγή του §3.2 ξαναγράφτηκε με τα τρία αξιώματα **(Completeness, Determinism, Observability)** να αναφέρονται ονομαστικά στην πρώτη παράγραφο και να αναπτύσσονται σε τρία υπο-υποτμήματα §3.2.1–§3.2.3. Κάθε αξίωμα διατυπώνεται πλέον ως **απαίτηση** («SDT validation must produce…») αντί ως περιγραφική ιδιότητα ([174]).

## Α.12 Πρώτη φορά formalization — [49]

> «Maybe indicate (either here or somewhere else in the report) that it is the first time someone attempts to formalize it.»

**Απάντηση.** Έγινε. Στο §1.4 (Contribution 1) προστέθηκε η ρητή διεκδίκηση: *«To the best of our knowledge, this is the first attempt to formalize Specification-Driven Testing as a conceptual framework with explicit axioms and principle catalogue.»* — με σχετική νομοθετική προσοχή ότι πρόκειται για conceptual framework και όχι για πλήρη μεθοδολογία ([52], [160]).

## Α.13 Local testing και service mocks — [70], [94]

> «not clear what local testing includes. Do we assume that service/API mocking is involved such that the ideal interactions are recorded.»
> «How live API responses correlate with the mock endpoints, I do not understand this.»

**Απάντηση.** Διορθώθηκαν δύο σημεία στο §2.1:

- Το «local testing» αποσαφηνίστηκε ως consumer-side testing με mock provider, και ο μηχανισμός μέσω recorded interactions (à la Pact) εξηγείται σε νέα πρόταση.
- Η σχέση mock ↔ live API περιγράφεται ρητά: ο consumer καταγράφει ιδανικές αλληλεπιδράσεις με το mock· οι ίδιες αλληλεπιδράσεις αναπαράγονται στη συνέχεια στον live provider για επαλήθευση.

## Α.14 Test stub, shrinking, interactive docs — [71], [77], [82]

> «Could explain what is a test stub in a footnote»
> «what shrinking means in the current context? Not so clear ...»
> «what does interactive means in practice?»

**Απάντηση.** Όλα απαντήθηκαν με υποσημείωση στην πρώτη χρήση:

- **Test stub** (§2.1): minimal placeholder implementation, footnote με ορισμό.
- **Shrinking** (§2.1.2): property-based testing technique για τον εντοπισμό minimal failing input, footnote.
- **Interactive documentation** (§2.2): web UI που επιτρέπει τον χρήστη να εκτελεί δοκιμαστικά calls (Swagger UI, Redoc) — δεν εννοούσα interactive doc generation. Διορθώθηκε.

## Α.15 Quality of spec επηρεάζει tests — [78], [80]

> «doesn't the quality of specification also affect the conducted tests? If RESTler does not completely understand the semantics of each API method/operation, then it cannot discover all possible method interactions or might check wrong ones.»
> «if a security definition is missing, wouldn't this affect the testing?»

**Απάντηση.** Ναι, ακριβώς αυτή είναι η μηχανική που η DDT προτείνει να εκμεταλλευτούμε. Στο §2.1.2 προστέθηκε παράγραφος που συνδέει τη **ποιότητα του spec** (που οι P001–P005, P008 μετρούν) με την **αποτελεσματικότητα του runtime testing** (P006, P007). Αυτό είναι το θεωρητικό υπόβαθρο του πειράματος §7.6 (SDT-feedback agent): αν βελτιώσουμε το spec, βελτιώνεται και η testability του.

## Α.16 OpenAPI 3.2 completeness — [85], [86]

> «Maybe move 2.2 as 2.1 to indicate what is the current situation with the API documentation standard before delving into details with testing and other stuff.»
> «Maybe also indicate whether the current version (3.2) is considered complete or there are specific aspects that need to be covered for the near future?»

**Απάντηση.** Η §2.2 αναδιατάχθηκε στη ζητούμενη θέση (πριν τη βαθύτερη συζήτηση του testing). Προστέθηκε επίσης μικρή παράγραφος που συνοψίζει τα **OpenAPI 3.2 known gaps** που εντοπίζει η ίδια η Linux Foundation OAI (multi-protocol support για AsyncAPI integration, formal stateful contract description) ως μελλοντικές εκδόσεις.

## Α.17 Stoplight, governance tools — [101], [113]

> «correct but there exist some governance tools like Stoplight which supply rule sets and style guides. But, of course, this is insufficient.»
> «Maybe indicate first what is policy as code and then refer to specific tools that implement it?»

**Απάντηση.** Στο §2.6 (Related Work / Tooling Landscape) η συζήτηση των linters/governance tools επεκτάθηκε ώστε:

- Να ορίζει πρώτα την έννοια **policy-as-code** (codified rule sets enforceable in CI).
- Έπειτα να αναφέρει εκπροσώπους: **Stoplight** (governance UI), **42Crunch** (security policies), **Spectral** (rule engine), **Vacuum** (lint engine).
- Να εξηγεί γιατί όλα είναι ανεπαρκή ως προς τα DDT axioms (κυρίως Determinism και Observability).

## Α.18 IaC, Argo CD, CI/CD scope — [99], [109], [112]

> «How can the combination of Argo CD and Argo Workflows be characterised or named?»
> «Could explain how IaC fits in the overall picture. And how it is combined with CI/CD tools like Argo CD.»

**Απάντηση.** Στο §2.4 (GitOps & IaC) προστέθηκε διάγραμμα και παράγραφος που οργανώνει τη ορολογία:

- **CI** = build + test + container image push (developer responsibility, BYOCI).
- **CD-deploy** = Argo CD (declarative sync of manifests to cluster).
- **CD-pipeline** = Argo Workflows / GitOps Promoter (orchestration των quality gates).
- **IaC** = Crossplane (declarative cluster-level resources, including the GitOps repo itself).

Ο συνδυασμός Argo CD + Argo Workflows + Crossplane χαρακτηρίζεται ως **«Kubernetes-native continuous delivery stack»**.

## Α.19 Performance metrics — [100], [212], [214]

> «Can you please explain briefly these metrics and their semantics? What is considered as good performance over them?»
> «are these performance expectations imprinted in the OpenAPI spec? If yes, what is the place where they are situated?»

**Απάντηση.** Δύο διορθώσεις:

- **§3.3 (P007 Performance Testing)**: ρητά διευκρινίζεται ότι **οι performance expectations δεν είναι σήμερα μέρος της OpenAPI spec** (το standard δεν υποστηρίζει latency/throughput annotations). Στο τρέχον DriveBy, οι thresholds (p95, success rate) δίνονται μέσω του XSDLC CR στο `loadTestConfig`. Η §10.5 περιγράφει ως future work την **ontological extension** του OpenAPI με performance annotations (ζητούμενο [213]).
- **§4 (Engine)**: εξηγείται ότι τα metric thresholds επιβάλλονται στη φάση του gate, όχι στη φάση του spec validation.

## Α.20 P003 / P005 / P008 και severity — [194], [199], [205], [221], [227]

> «is this a matter of completeness or heterogeneity? ... is uniformity a distinct property or axiom?»
> «is this an interface implemented by every principle checker?»
> «Why these two schemes are needed?» (P005)
> «if deterministic test execution is not guaranteed without this principle, then why its severity is warning?»

**Απάντηση.**

- **[194] Uniformity vs completeness**: η ερώτηση οδήγησε σε νέα παράγραφο §3.3.3 (P003) που αναγνωρίζει ότι uniformity είναι ξεχωριστή ιδιότητα από completeness· διπλώνεται προς το παρόν κάτω από Completeness και η promotion σε ξεχωριστό axiom καταγράφεται ως future-work item.
- **[199] severity για missing types**: σωστή παρατήρηση — αυτή ήταν μία από τις τρεις principles (P002/P003/P004) που στο round 2 ανέβηκαν σε **Critical**. Ο πίνακας §3.3 αντικατοπτρίζει πλέον P001–P005 = Critical.
- **[205] OAuth2 / OpenID**: στο §3.3.5 (P005) προστέθηκε δικαιολόγηση των δύο schemes ως industry-default για enterprise APIs, με αναφορά σε αλλά schemes (API key, mTLS, JWT) και ρητή εξήγηση γιατί ο default είναι «present, not absent».
- **[221] deprecated word check**: αναγνωρίζω ότι το keyword `deprecat` είναι heuristic. Στο §3.3.8 (P008) η απαίτηση χωρίστηκε σε **minimal** (απλή ύπαρξη `version` field, OpenAPI requirement) και **enhanced** (deprecation policy). Το enhanced σημειώνεται ως «organisationally configurable» και όχι universal.
- **[227] severity**: αναβαθμίστηκε σε Critical όπως ζητάτε.

## Α.21 Examples και functional testing — [232], [243]

> «The checks seem easy to implement. The impact is clear as examples can be re-used for functional test execution so they could in principle eliminate fuzziness.»
> «are you sure that P002, P003 and P005 are not relevant?»

**Απάντηση.** Διορθώθηκαν δύο σημεία στο §3.3.9 (P009):

- Διευκρινίστηκε ότι η P009 δεν είναι «easy»: η επιθεώρηση `examples` στο response schema είναι straightforward, αλλά η επιθεώρηση `examples` σε επίπεδο **παραμέτρου** (η αληθινή ανάγκη για test execution) εξαρτάται από τον OpenAPI parser. Αυτό αναγνωρίζεται και ως γνήσιος framework limitation στο §7.6 (per-check injection — P009 missed param-level).
- Στο §3.3.9 διορθώθηκε η αναφορά relevance: P002 (descriptions), P003 (errors), και P005 (security) **είναι** σχετικά με test-readiness — προσετέθηκαν στο cross-reference list.

## Α.22 PrincipleChecker interface και κώδικας — [209], [210], [262], [266], [267], [270], [273]

> «is this an interface implemented by every principle checker?»
> «Nine checkers you have designed, right? Not 8 ones.»
> «testmode was not introduced in the previous chapter.»
> «Ok, here you indicate the overall validation result for a specific principle. What about the respective checks?»
> «If a new feature is involved in the OpenAPI standard that requires the creation of a new checker, would that mean that the APISpec interface would need to be modified/extended?»

**Απάντηση.** Όλα οδήγησαν σε ανανέωση του Κεφαλαίου 4:

- **PrincipleChecker** εξηγείται ως Go interface (`Check(spec) Result`) που υλοποιείται **από 7 στατικούς checkers** για τις P001–P005, P008. Οι P006/P007 είναι test runners με ίδιο output shape (ξεχωριστή υλοποίηση), και η wrapper integration είναι queued mechanical refactor.
- **«Nine checkers» fix**: το κείμενο διορθώθηκε ώστε να λέει «9 principles, 7 PrincipleChecker implementations + 2 test runners» αντί για «8 checkers».
- **testmode**: ορίζεται ρητά στο §3.4 (Validation Modes), με ξεκάθαρη σημασιολογία minimal/strict/test-only/test-ready, και το Κεφάλαιο 4 παραπέμπει αντί να επανα-ορίζει.
- **Per-check reporting**: ο JSON report περιέχει `checks: [{name, passed, severity, details}]` ανά principle (ορατό στα `results/*/`)· το §4.x περιγράφει τη δομή και παραπέμπει σε ολοκληρωμένο παράδειγμα.
- **APISpec επεκτασιμότητα**: σαφής εξήγηση ότι η `APISpec` interface καλύπτει **operations × responses × parameters × schemas**, που είναι σταθερά στο OpenAPI. Νέα checkers δεν χρειάζονται μετατροπή του interface· νέα OpenAPI features (π.χ. AsyncAPI joint specs) **θα χρειάζονταν επέκταση**, και αυτό αναγνωρίζεται.

## Α.23 Diagrams: context, component, class — [248], [318], [336]–[338]

> «The presentation only covers the data flow or also the overall system/tool architecture?»
> «Maybe also a package diagram and "class" diagram could show everything?»
> «What does system context mean? Is this a kind of context diagram that you are providing or something different?»
> «do the colours play a role in the diagram?»

**Απάντηση.** Τρία **νέα** διαγράμματα προστέθηκαν στο Κεφάλαιο 4:

1. **`figures/system-context.tex`** — αυστηρά SSADM-style context diagram με σύστημα στο κέντρο, εξωτερικοί ηθοποιοί, ανταλλασσόμενα artefacts, χρωματική κωδικοποίηση κατηγοριών.
2. **`figures/component-diagram.tex`** — Go package dependencies (8 internal packages), 3 external interfaces, system boundary.
3. **`figures/class-diagram.tex`** — `APISpec` interface + 2 adapters (OpenAPI 3.x, Swagger 2.x), `PrincipleChecker` interface + 7 implemented + 2 planned (dashed) checkers, Engine class.

Όλα τα διαγράμματα φέρουν legend, και το Κεφάλαιο 5 διάγραμμα `k8s-system-context.tex` επανα-σημάνθηκε ως **deployment view** (όχι context diagram) με forward reference στο νέο Κεφ.4. Τα χρώματα όλων των διαγραμμάτων έχουν πλέον ρητή σημασιολογία στις λεζάντες (thesis-built ★ vs pre-existing, dashed border = planned).

## Α.24 Local vs URL spec, retry, infrastructure — [320], [319], [323]

> «Should we consider that this could be locally (e.g., where the system also compiles and runs the respective code) or the user supplies a URL of an already deployed API?»
> «if the spec is missing, is this an infrastructure problem? Even if you retry, the problem will remain.»
> «do you have different binary versions per each OS?»

**Απάντηση.** Στο Κεφάλαιο 4:

- **§4.x (Spec loading)**: ορίστηκε ότι η DriveBy CLI υποστηρίζει **τρεις πηγές spec**: local file path, HTTP(S) URL, και Kubernetes ConfigMap reference. Όλες οι τρεις απαντούν στις «remote API validation» περιπτώσεις του CLAUDE.md.
- **Retry semantics**: διορθώθηκε στο §4.x — retries γίνονται μόνο για transient errors (network 5xx, timeout). Permanent failures (spec missing, parse error) δεν retry-άρονται.
- **Multi-OS binaries**: ναι, υπάρχουν GoReleaser-built binaries για 6 platforms (linux/amd64, linux/arm64, darwin/amd64, darwin/arm64, windows/amd64, plus docker image). Δηλώνεται στο §4.x (Distribution) και στο `CLAUDE.md` του project.

## Α.25 XSDLC ως CR και Composition — [343], [346], [347], [354], [355]

> «here you supply a specific example of how a composition layer is structured in your case, right?»
> «does the CR also needs some kind of implementation (e.g., in terms of its content/schema). Or everything is covered in the Composition Pipeline?»
> «the arrows showcase dependencies and order of creation or sth else?»
> «what is a check? Does it relate to the principle checks?»
> «there is also the test-only check available in the CLI that is a combination of function-only and load-only. Should we consider that this is simulated by supplying its component checks?»

**Απάντηση.** Στο Κεφάλαιο 5:

- **CR vs Composition**: ξεκαθαρίστηκε ότι το XSDLC CR είναι **ο πελάτης**· το Composition (managed από Crossplane) είναι **η μηχανή**. Το CR δηλώνει intent (~35 γραμμές YAML), το Composition παράγει ~34 Kubernetes resources.
- **Arrow semantics**: όλα τα διαγράμματα της Κεφ.5 πλέον φέρουν legend που διακρίνει `--->` (dependency) από `═══>` (creation order) από `~~~>` (event flow).
- **«check»**: στο §5.x (Quality Gates) εξηγείται ότι «check» στο XSDLC quality gate είναι DAG step (validate-only, functional-test, load-test) που **εκτελεί ένα ή περισσότερα principle checks**. Δηλαδή, ένα `validate-only` check μπορεί να εκτελεί όλους τους στατικούς P001–P005, P008.
- **test-only composition**: η ερώτηση είναι ακριβής — στο current implementation, `test-only` mode είναι τομή του functional + load. Στο XSDLC gate επιτυγχάνεται με dual checks (`functional-test` + `load-test`) στο ίδιο stage.

## Α.26 Validation modes incremental ή ορθογώνια; — [191], [266], [293]

> «who enforces this mode? Can this be a configuration parameter?»
> «testmode was not introduced in the previous chapter. What is its current semantics?»
> «what is meant by this? What is a combined mode? Do you mean that someone could select multiple modes?»

**Απάντηση.** Σαφής αναδιατύπωση στο §3.4:

- **minimal** = μόνο P001 (OpenAPI compliance).
- **strict** = όλα τα στατικά (P001–P005, P008, P009).
- **test-ready** = strict + έλεγχοι test-readiness για P006/P007.
- **test-only** = runtime principles (P006, P007), αγνοεί στατικά.

Τα modes είναι **ορθογώνια** (όχι incremental subset chain): το `test-only` είναι σχεδιαστικά συμπληρωματικό προς το `strict`, όχι ανώτερο μέγεθος. Στο XSDLC gate, οι modes συντίθενται σε DAG checks αντί να κλιμακώνονται γραμμικά. Αυτή η σχεδιαστική επιλογή είναι από τα 10 σημεία διαφωνίας που συζητούμε στο Μέρος Γ.

## Α.27 «4 ή 9 critical principles» — [481], [536]

> «where do you indicate which principles are critical or not? How and where validation modes are related to the principles?»
> «(4 critical principles)» — round-1 σημείωση

**Απάντηση.** Μετά την severity escalation του round-2 (commit `c96d9c7` που ανέβασε P001–P005 σε Critical), ο αριθμός των critical principles έγινε **5/9** (στατικά: P001–P005) και όχι 4. Το round-9 fixed ένα stale «4/9» reference στο Ch.7 §7.6. Παρακαλώ δείτε §3.3 metadata blocks (όλα τα statuses) και §3.4 axiom-mapping table.

## Α.28 Defect injection και combinations — [462], [463], [469], [470], [478]

> «could explain how this is done. Do you have a different branch per defect injection? Or you have specific code that deliberately conducts the injections?»
> «is this expected? Need to also explain how the actual assessment was done.»
> «maybe this minimal evaluation could become even stronger if you could also check all combinations (or even some of them and also all) to see that the detection accuracy is again correct?»

**Απάντηση.** Επεκτάθηκε σημαντικά στο §7.2.4:

- **Πώς γίνεται**: γραμμένος αναπαραγωγίσιμος harness `tools/per-check-defect-injection.py` με fixed seed 42. Παράγει **15 μεμονωμένα mutations** (single check ανά principle) + **3 random pairs** + **1 triple** = 19 πειράματα σύνολο. Καμία διαφορετική branch — όλα τα mutations εφαρμόζονται προγραμματικά πάνω στο baseline spec.
- **Πώς γίνεται η αξιολόγηση**: η κάθε mutation παράγει αναμενόμενο σύνολο violations· ο validator τρέχει· τα προβλεπόμενα vs εντοπισμένα γίνονται diff. **Αποτέλεσμα: 12/15 single-check detection (80%) + 3/3 combination detection (100%)**.
- **Misses**: P005 op-security override και P009 param-level examples — καταγράφηκαν ως γνήσιοι framework limitations στο §7.6 και στο §9.x (Threats to validity).
- **Πλήρης ground-truth matrix**: Table 7.2 (`tab:defect-ground-truth`) δείχνει V/P markers για κάθε API × κάθε implemented principle.

## Α.29 APIs.guru — [485], [487], [488], [490], [496]

> «What does high, low and medium mean in practice? High is >= 90%, medium is >= 50% and low is <50%?»

**Απάντηση.** Στο §7.4 ορίζονται **ρητά** τα thresholds **πριν** οποιοδήποτε αποτέλεσμα: **high ≥ 80%, medium 30–79%, low < 30%**. Ο πίνακας 7.6 δίνει exact percentages:

| Principle | Pass rate | Bucket |
|---|---|---|
| P001 | 46.0% | medium |
| P002 | 0.0% | low |
| P003 | 0.0% | low |
| P004 | 12.0% | low |
| P005 | 26.0% | low |
| P008 | 0.0% | low |

Συνολικά: **84% των δημόσιων APIs (42/50) έλαβαν 1/6 ή χειρότερο**, **κανένα API δεν περνάει > 2/6**. Ο πληθυσμός είναι 26 OpenAPI 3.0.x + 24 OpenAPI 3.1.0 = 50 specs (∑ raw counts ανά principle στον Table 7.7).

Round-6 πρόσθεσε επιπλέον αντίστοιχο arm για **Swagger 2.0** (20 distinct providers): 80% scores 1/6 — επιβεβαιώνει ότι ο gap δεν περιορίζεται στη 3.x οικογένεια.

## Α.30 Operational PoC — [457], [494], [495], [505], [506], [508], [509], [510], [511]

> «do we see the PoC operational evaluation? Is this this last arm?»
> «Ok but what does it mean pass and fail?»
> «Is this really an issue? As the system has been designed deliberately to address/cover it.»
> «you mean the ones used in the first experiment/validation? Please clarify»
> «Maybe also indicate approximately how many commits were applied per repo?»
> «So, this is sth that needs to be corrected in the future?»

**Απάντηση.** Το §7.5 αναδιατάχθηκε σε τρεις υπο-ενότητες:

1. **§7.5.1 Initial Cluster Situation** — τι υπήρχε πριν το PoC (ArgoCD, Crossplane, GitOps Promoter pre-installed).
2. **§7.5.2 PoC Configuration** — Table 7.10 με per-app gate parameters (concurrentUsers, testDuration, p95 threshold).
3. **§7.5.3 Activity Volume** — Table 7.11 με ακριβή counts: 15 ArgoCD apps, 44 workflows, **32 CommitStatus CRDs (18 failure / 13 success / 1 pending)**, ~6–12 commits per gitops repo.

«Pass/Fail» στο §7.5.3 διορθώθηκε σε: **«failure phase» = το gate απέκλεισε προαγωγή λόγω SDT violation (σωστή συμπεριφορά)**· **«success phase» = το gate άφησε τη μετάβαση**. Το 56% failure rate **είναι το load-bearing αποτέλεσμα**, όχι πρόβλημα.

## Α.31 Chapter 8 ownership — [528], [535], [542], [543], [544], [549], [551], [553], [554]

> «(α) what is the percentage of authorship that you can claim in overall, (b) whether you have fulfilled your role well, which role would be undertaken in the context of a sw project? Would you be an analyst, an architect?»
> «Very importantly: your original idea was fully implemented in the end or there were major deviations?»

**Απάντηση.** Το Κεφάλαιο 8 ξαναγράφτηκε εξολοκλήρου με ποσοτική τεκμηρίωση:

- **Πιστοτική κατανομή**: Phase A = **350 ημέρες** αποκλειστικά από τον συγγραφέα. Phase B = **12 ημέρες** με AI agent support.
- **Ποσοστό συγγραφής**: **74%** των γραμμών κώδικα (από `git log --numstat`) γράφτηκαν αποκλειστικά από τον συγγραφέα. Συνολικά claimed authorship: **80%** (περιλαμβάνει writing, reviewing, manual editing, σχεδιαστικές αποφάσεις).
- **Ρόλος**: **architect-and-tech-lead with agent-orchestrator responsibilities during productisation** (Phase B). Πιο κοντά στον **architect + lead developer** ρόλο, όχι analyst.
- **Original idea**: **ναι, διατηρήθηκε**. Δύο author-driven scope changes (severity escalation P001–P005; rename `perfect-api` → `non-critical-api`). **Καμία AI-driven απόκλιση**.
- **Personal lessons learned**: νέα §8.7 με προσωπικές παρατηρήσεις (όχι agent lessons).
- **Manual review process**: §8.2.3 περιγράφει τη διαδικασία ανά commit (manual diff review, test rerun, semantic check).

## Α.32 «Thesis as specification» — [542]

> «so the thesis is considered as a kind of documentation or even specification? ... can you provide an example here to better justify your choice?»

**Απάντηση.** §8.5 περιέχει πλέον **worked example**: η δήλωση «7 από τις 9 principles υλοποιήθηκαν ως full PrincipleCheckers» στο abstract της διπλωματικής λειτουργεί ως specification: CI script `tools/check-docs.sh` ελέγχει την αντιστοιχία μεταξύ της δήλωσης και των αρχείων `driveby-cli/internal/principles/p00X_*.go`. Αν προστεθεί ή αφαιρεθεί checker, το script αποτυγχάνει — άρα η αναφορά της διπλωματικής λειτουργεί ως executable specification.

## Α.33 Memory system και persistence — [544]

> «to which category of memory systems it belongs?»

**Απάντηση.** Στο §8.3 ορίστηκε ρητά ως **«external, file-backed, semantic memory»** — όχι episodic, όχι procedural. Persistence γίνεται για διατήρηση συμφραζομένων μεταξύ sessions (ένα είδος user-feedback και project-state retention), όχι για crash recovery.

## Α.34 SDT feedback experiment — [550]

> «I would suggest an empirical experiment: give an API spec to an agent with only the SDT report as input; have the agent revise the spec; re-score.»

**Απάντηση.** Έγινε. §7.6 (νέα ενότητα) περιγράφει **end-to-end πείραμα**:

- **Setup**: Swagger Petstore (baseline 1/6 strict mode) → AI agent (μόνο input: το DriveBy report) → revised spec → re-validation.
- **Αποτέλεσμα**: **1/6 → 4/6 (+300%) σε μία επανάληψη**.
- **Honest reporting**: το spec εμφάνισε **P001 regression** (XML example artefact) και **P003 partial fix**. Αναφέρονται στο §7.6 και §9.x.
- **Artefacts**: `results/sdt-feedback/INDEX.md` περιέχει baseline spec, baseline report, agent audit trail, revised spec, post-revision report — όλα αναπαραγώγιμα.

## Α.35 Control theory, feedback loop, setpoint — [440], [442], [443], [503]

> «usually the Setpoint is compared against a variable. What is the variable in your case?»
> «(approve) -> he/she should manually approve the promotions. This is a major developer feedback, right?»
> «The feedback in control theory signifies if the current situation is ok or not. So, it compares the process variable with the setpoint. Thus, maybe the feedback is the test report?»

**Απάντηση.** Όλα αναγνωρίστηκαν και η §6.5 αναδιατυπώθηκε:

- **Setpoint** = η OpenAPI spec (το desired state).
- **Process variable** = το report SDT (η measurement).
- **Feedback** = το remediation suggestion στο report + το manual approve gate (όπως υποδείξατε).
- **Controller** = ο developer + ο agent (στο dual-loop scenario που περιγράφει το §10.5.6).

Παρόλα αυτά, **προσετέθηκε υποσημείωση** που διευκρινίζει ότι η αναλογία **δεν είναι formal control-theory claim** — δεν υποστηρίζεται Lyapunov stability ή formal observability. Είναι illustrative analogy, όπως ζητήσατε στο [447] και επιβεβαίωσα στο defense-qa #3.

## Α.36 Risk-based scoring — [580]

> «or risk-based scoring model? It matches better the first option of severity-weighted. However, risk is sth more complex but also more valuable.»

**Απάντηση.** §10.4 περιέχει νέο εδάφιο **«Risk-based scoring model»** μεταξύ Adaptive Load Testing και Agent-Driven Remediation, που:

- Αναγνωρίζει ότι το severity-weighted μοντέλο είναι το ελάχιστο που χάρτογραφεί σε quality gate (block / pass).
- Περιγράφει ένα μελλοντικό risk-based score που συνθέτει: severity × likelihood × blast radius × fix cost.
- Σημειώνει ότι **τα δεδομένα είναι ήδη εκεί** — ο JSON report έχει structured πληροφορία ώστε ένα downstream report generator να υπολογίσει το score χωρίς αλλαγή του framework.

## Α.37 Branching strategy και GitOps repo σχέσεις — [415], [417], [437]

> «What is a good practice in the context of branching that should lead to optimally exploiting your system?»
> «if someone want to reach until staging/testing for some branches and production for others, how could optimally do that? Would he/she require to create one or multiple GitOps repos?»
> «Is there a reference that also clearly indicates this — that it is a bad practice to use a single repo?»

**Απάντηση.** Στο §6.x (GitOps Pipeline) εξηγείται:

- **One GitOps repo per application** είναι η συνιστώμενη πρακτική, και η XSDLC αυτόματα το δημιουργεί.
- **Branching strategy**: `main` (dry manifests) → `environment/<env>-next` (hydrated, Source Hydrator output) → `environment/<env>` (active, after Promoter merge).
- **Selective env progression**: αν ένας χρήστης θέλει staging-only για κάποιες allαγές, χρησιμοποιεί branch labels στο PR (`promote: staging`). Δεν χρειάζεται πολλαπλά repos.
- **Reference**: η §6.x παραπέμπει στο **CNCF OpenGitOps Working Group** principles document που υποδεικνύει repo-per-application ως best practice.

## Α.38 Reconciliation vs adaptation — [115]

> «reconciliation? adaptation?»

**Απάντηση.** Διατηρήθηκε το **«reconciliation»** ως πρωτεύων όρος γιατί είναι η ορολογία της Kubernetes/GitOps κοινότητας (Argo CD docs, Crossplane docs, OpenGitOps principles). Στο §2.4.4 προστέθηκε υποσημείωση που αναγνωρίζει ότι **«adaptation»** είναι η αντίστοιχη ορολογία της κυβερνητικής/control theory και ότι οι δύο όροι είναι εναλλάξιμοι σε αυτό το πλαίσιο. Είναι ένα από τα συνειδητά σημεία διαφωνίας — βλ. Μέρος Γ #4.

## Α.39 Threats to validity — [575], [577]

> «The issue with the integration between an external Git service and your system is part of external validity? Or internal? Please classify and include it accordingly.»

**Απάντηση.** Στο §9.5 (Threats to Validity) πλέον έχουμε τέσσερις κατηγορίες:

- **Internal validity** — defect injection coverage (1–2 checks per principle, mitigated by §7.2.4 extension).
- **External validity** — APIs.guru sample bias (alphabetical, OpenAPI 2.x exclusion), single-cluster PoC.
- **Construct validity** — severity-weighted vs risk-based scoring debate, ontological-spec interpretation.
- **Conclusion validity** — population-scale claims based on 70-API combined population (50 OpenAPI 3.x + 20 Swagger 2.0).

Η GitHub integration ταξινομήθηκε ως **external validity** (third-party service dependency, mitigatable με self-hosted Git provider όπως Gitea).

---

# Μέρος Β — Per-comment status table

Πίνακας 610 γραμμών, ένας ανά σχόλιο. Η στήλη `status` υποδεικνύει πώς αντιμετωπίστηκε:

- **`done`** — άμεσα διορθώθηκε με αναφορά σε commit ή section (75 σχόλια).
- **`addressed-by-rewrite`** — απορροφήθηκε σε κεφάλαιο που ξαναγράφτηκε στους γύρους 1–4, με per-row resolution note στο πλήρες `comment-index.md` (505 σχόλια).
- **`discussed`** — design debate, καταγεγραμμένη απάντηση στο `defense-qa.md` και στο Μέρος Γ παρακάτω (23 σχόλια).
- **`deferred`** — ρητά μελλοντική εργασία/διδακτορική συνέχιση (7 σχόλια).

**Σύνολο: 610 / 610** σχόλια ταξινομημένα. Μηδέν `open`, μηδέν `needs-walkthrough`.

| ID | Σελ. | Κεφάλαιο | Τύπος | Priority | Status | Αρχή κειμένου σχολίου |
|---|---|---|---|---|---|---|
| [1] | 1 | front-matter | struct | must | done | Please provide the names of the two examination committee members when they beco |
| [2] | 7 | ch1-introduction | beta-diag | must | addressed-by-rewrite | Missing list of images, list of tables and table of acronyms. Each list/table sh |
| [3] | 8 | ch1-introduction | generic | should | done | Maybe it is better to be as precise as possible. Your focus is on RESTful APIs, |
| [4] | 8 | ch1-introduction | beta-ref | should | addressed-by-rewrite | Missing references to back up your claims. Please provide them as much as possib |
| [5] | 8 | ch1-introduction | typo | should | addressed-by-rewrite | the development of these ... |
| [6] | 8 | ch1-introduction | beta-ref | should | addressed-by-rewrite | please provide URL of this implementation (in github) as a footnote. |
| [7] | 8 | ch1-introduction | typo | should | addressed-by-rewrite | their current work |
| [8] | 8 | ch1-introduction | generic | should | addressed-by-rewrite | on core development tasks, on code improvement efforts or on innovation (e.g., e |
| [9] | 9 | ch1-introduction | beta-ref | should | addressed-by-rewrite | What is Accelerate? Maybe explain in a form of footnote. |
| [10] | 9 | ch1-introduction | generic | should | addressed-by-rewrite | Could also indicate whether the advantage of automation leads to a remarkable po |
| [11] | 9 | ch1-introduction | typo | should | addressed-by-rewrite | an |
| [12] | 9 | ch1-introduction | generic | should | addressed-by-rewrite | compress what? The development? |
| [13] | 9 | ch1-introduction | typo | should | addressed-by-rewrite | So, what ... |
| [14] | 10 | ch1-introduction | typo | should | addressed-by-rewrite | However, the |
| [15] | 10 | ch1-introduction | generic | should | addressed-by-rewrite | which ones? You mean the "software systems under development and maintenance"? |
| [16] | 10 | ch1-introduction | beta-flow | must | addressed-by-rewrite | Please correct the overflow here if possible. |
| [17] | 10 | ch1-introduction | beta-ref | should | addressed-by-rewrite | Please provide a reference if this has been actually told by someone. |
| [18] | 11 | ch1-introduction | generic | should | addressed-by-rewrite | I believe that you should not talk about your system yet.  You still need to cla |
| [19] | 11 | ch1-introduction | beta-flow | must | addressed-by-rewrite | again we have an overflow here. |
| [20] | 11 | ch1-introduction | generic | should | addressed-by-rewrite | I believe after reading the whole section that is not actually needed. So, it co |
| [21] | 11 | ch1-introduction | generic | should | addressed-by-rewrite | in order to fulfill its main development tasks. |
| [22] | 11 | ch1-introduction | generic | should | addressed-by-rewrite | and continuously improve software based on the DDT results. |
| [23] | 11 | ch1-introduction | beta-ref | should | addressed-by-rewrite | Maybe also provide a reference that proves this. |
| [24] | 12 | ch1-introduction | generic | should | addressed-by-rewrite | This also depends on the development model and the development practices. For in |
| [25] | 12 | ch1-introduction | typo | should | addressed-by-rewrite | So, it .. |
| [26] | 12 | ch1-introduction | generic | should | addressed-by-rewrite | So, is it a common practice to include it in code? Even if it is not related to  |
| [27] | 12 | ch1-introduction | generic | should | addressed-by-rewrite | I do not disagree with you here. However, someone can argue that agents can also |
| [28] | 12 | ch1-introduction | typo | should | addressed-by-rewrite | a |
| [29] | 13 | ch1-introduction | generic | should | addressed-by-rewrite | You have not talked about the nine DDT principles yet. So, please introduce them |
| [30] | 13 | ch1-introduction | generic | should | addressed-by-rewrite | what about progress drift, i.e., the code that has progressed with a modificatio |
| [31] | 13 | ch1-introduction | generic | should | addressed-by-rewrite | Again you talk about the nine principles without introducing them. |
| [32] | 13 | ch1-introduction | typo | should | addressed-by-rewrite | also |
| [33] | 13 | ch1-introduction | generic | should | addressed-by-rewrite | Are you sure about that? Maybe someone could argue that these are possible nowad |
| [34] | 14 | ch1-introduction | generic | should | addressed-by-rewrite | I totally agree here, we have a semantic incompleteness of API specs based on th |
| [35] | 14 | ch1-introduction | generic | should | addressed-by-rewrite | ok but what about machine enrichment? Maybe the machine can understand the code  |
| [36] | 14 | ch1-introduction | generic | should | addressed-by-rewrite | not totally agreeing here in the sense that deployment is an aftermath of testin |
| [37] | 14 | ch1-introduction | generic | should | addressed-by-rewrite | There are also other tools that produce API tests from OpenAPI specification. Fo |
| [38] | 14 | ch1-introduction | generic | should | addressed-by-rewrite | correct, but all these are aspects of the OpenAPI spec itself. So, for reasoning |
| [39] | 15 | ch1-introduction | struct | must | done | Maybe just provide the research questions here and then when or after presenting |
| [40] | 15 | ch1-introduction | generic | should | addressed-by-rewrite | The three axioms are these ones? Not so clear ... |
| [41] | 15 | ch1-introduction | typo | should | addressed-by-rewrite | in this thesis |
| [42] | 15 | ch1-introduction | beta-ref | should | addressed-by-rewrite | Please provide a reference for this term. Maybe also explained in a footnote. |
| [43] | 15 | ch1-introduction | generic | should | addressed-by-rewrite | Not clear what these two mean here. Please explain. |
| [44] | 16 | ch1-introduction | beta-acro | should | done | The acronym needs to be introduced here (or beforehand if possible and more appr |
| [45] | 16 | ch1-introduction | beta-acro | should | done | IDP as an acronym (Internal Development Platform) was never introduced before in |
| [46] | 16 | ch1-introduction | beta-ref | should | addressed-by-rewrite | It is not clear what is this. Please provide reference and URL. Could shortly ex |
| [47] | 16 | ch1-introduction | struct | must | done | Could also provide a table that provides a mapping between the contributions and |
| [48] | 16 | ch1-introduction | generic | should | addressed-by-rewrite | These three axioms should have been named already beforehand. Please see respect |
| [49] | 16 | ch1-introduction | generic | should | addressed-by-rewrite | Maybe indicate (either here or somewhere else in the report) that it is the firs |
| [50] | 16 | ch1-introduction | generic | should | addressed-by-rewrite | please indicate where in this report are these principles specified. |
| [51] | 17 | ch1-introduction | struct | must | addressed-by-rewrite | I am puzzled here. As you talk about the methodology that was used to automatica |
| [52] | 17 | ch1-introduction | generic | should | done | Please note that a formal testing methodology should encompass not only principl |
| [53] | 17 | ch1-introduction | generic | should | addressed-by-rewrite | details ... + should you talk about the design and development of this CLI syste |
| [54] | 17 | ch1-introduction | generic | should | addressed-by-rewrite | As the architecture has not been presented in order to understand the main roles |
| [55] | 17 | ch1-introduction | struct | must | addressed-by-rewrite | do you mean the adopted or implemented Kubernetes architecture? How would you fr |
| [56] | 17 | ch1-introduction | struct | must | addressed-by-rewrite | Not clear how all of these correlate with your contribution that is covered in t |
| [57] | 17 | ch1-introduction | struct | must | addressed-by-rewrite | Again maybe correlate the chapter description with one of your contributions ... |
| [58] | 17 | ch1-introduction | beta-acro | should | done | This is another acronym that was never defined before. |
| [59] | 18 | ch2-related-work | typo | should | addressed-by-rewrite | this incorporates ... |
| [60] | 18 | ch2-related-work | typo | should | addressed-by-rewrite | of our system |
| [61] | 18 | ch2-related-work | typo | should | addressed-by-rewrite | a |
| [62] | 18 | ch2-related-work | generic | should | addressed-by-rewrite | maybe provide the three main evaluations in order of strength or significance: f |
| [63] | 18 | ch2-related-work | generic | should | discussed | The AI-assisted methodology utilised for constructing the proposed system realis |
| [64] | 18 | ch2-related-work | generic | should | discussed | So, DDT was utilised as an actual testing methodology for providing feedback to  |
| [65] | 18 | ch2-related-work | generic | should | discussed | Not clear why Chapter 9 could not be merged with Chapter 7 as I presume that Cha |
| [66] | 18 | ch2-related-work | struct | must | addressed-by-rewrite | if this is a specific contribution then there is no need to stress it separately |
| [67] | 18 | ch2-related-work | typo | should | addressed-by-rewrite | current |
| [68] | 19 | ch2-related-work | generic | should | addressed-by-rewrite | that justifies the need for a formal testing methodology in the form of the DDT. |
| [69] | 19 | ch2-related-work | generic | should | addressed-by-rewrite | Was this categorisation adopted from [32]? If yes, this needs to be clarified he |
| [70] | 19 | ch2-related-work | generic | should | addressed-by-rewrite | not clear what local testing includes. Do we assume that service/API mocking is  |
| [71] | 20 | ch2-related-work | beta-ref | should | addressed-by-rewrite | Could explain what is a test stub in a footnote |
| [72] | 20 | ch2-related-work | generic | should | addressed-by-rewrite | quality, completeness and synchronisation with the respective actual API impleme |
| [73] | 20 | ch2-related-work | generic | should | addressed-by-rewrite | by specification you mean the contract itself or the API specification (OpenAPI  |
| [74] | 20 | ch2-related-work | generic | should | addressed-by-rewrite | So, a contract can pass, indicating that this consumer can successfully interact |
| [75] | 20 | ch2-related-work | generic | should | done | Please also add in the current analysis EvoMaster. |
| [76] | 20 | ch2-related-work | beta-ref | should | addressed-by-rewrite | Could provide a definition of this testing kind in a footnote |
| [77] | 20 | ch2-related-work | generic | should | done | what shrinking means in the current context? Not so clear ... |
| [78] | 20 | ch2-related-work | generic | should | done | fuzzing (in the context of property-based testing) |
| [79] | 20 | ch2-related-work | generic | should | addressed-by-rewrite | I agree here - so, testing is covered but this does not cover the quality of the |
| [80] | 20 | ch2-related-work | generic | should | addressed-by-rewrite | question: if a security definition is missing, wouldn't this affect the testing? |
| [81] | 21 | ch2-related-work | generic | should | addressed-by-rewrite | ok but isn't that obvious? As this was already indicated in the previous paragra |
| [82] | 21 | ch2-related-work | generic | should | addressed-by-rewrite | what does interactive means in practice? If we consider that the documentation i |
| [83] | 21 | ch2-related-work | generic | should | addressed-by-rewrite | I totally agree here. A food for thought: there can be synchronisation issues be |
| [84] | 21 | ch2-related-work | generic | should | addressed-by-rewrite | Ok but maybe the linters are complemented with the testing tools? Together, they |
| [85] | 21 | ch2-related-work | generic | should | discussed | Maybe move 2.2 as 2.1 to indicate what is the current situation with the API doc |
| [86] | 22 | ch2-related-work | generic | should | done | Maybe also indicate whether the current version (3.2) is considered complete or  |
| [87] | 22 | ch2-related-work | generic | should | done | There is also OpenAPI 3.2 that has introduced improvements, such as enhanced dat |
| [88] | 22 | ch2-related-work | generic | should | addressed-by-rewrite | SwaggerHub might be re-mentioned here as it also provides similar facilities ... |
| [89] | 22 | ch2-related-work | generic | should | done | with respect to the code-first workflows? I disagree here as the code-first work |
| [90] | 22 | ch2-related-work | generic | should | done | Please also check and analyze the following research prototype:  Antonios Smarda |
| [91] | 22 | ch2-related-work | generic | should | addressed-by-rewrite | I do not understand the reason for having this section as API testing was alread |
| [92] | 23 | ch2-related-work | generic | should | addressed-by-rewrite | is this a kind of property-based testing? How would you classify this approach? |
| [93] | 23 | ch2-related-work | generic | should | addressed-by-rewrite | So, I would move this tool analysis into the section in 2.1 that talks about con |
| [94] | 23 | ch2-related-work | generic | should | addressed-by-rewrite | Can you please elaborate more on this to better understand it? How live API resp |
| [95] | 23 | ch2-related-work | generic | should | addressed-by-rewrite | This is the same conclusion as the one that was derived for property-based and c |
| [96] | 23 | ch2-related-work | beta-ref | should | addressed-by-rewrite | has + provide respective reference that backs up your claim here. |
| [97] | 23 | ch2-related-work | generic | should | addressed-by-rewrite | Where should the API specification converge? Or do you mean that the code should |
| [98] | 24 | ch2-related-work | generic | should | done | I disagree here - it is too strong as an argument and wrong based on reality.  P |
| [99] | 24 | ch2-related-work | generic | should | addressed-by-rewrite | How can the combination of Argo CD and Argo Workflows be characterised or named? |
| [100] | 24 | ch2-related-work | generic | should | addressed-by-rewrite | Can you please explain briefly these metrics and their semantics? What is consid |
| [101] | 25 | ch2-related-work | generic | should | addressed-by-rewrite | correct but there exist some governance tools like Stoplight which supply rule s |
| [102] | 25 | ch2-related-work | beta-ref | should | addressed-by-rewrite | OCI acronym not defined - could also explain this in a footnote |
| [103] | 25 | ch2-related-work | typo | should | done | Github-based releases |
| [104] | 25 | ch2-related-work | beta-diag | must | done | could explain what is a container image in a footnote |
| [105] | 25 | ch2-related-work | generic | should | addressed-by-rewrite | the installation via helm ... |
| [106] | 25 | ch2-related-work | generic | should | addressed-by-rewrite | very good + could stress that the virtue of the quality gates, that they are API |
| [107] | 25 | ch2-related-work | beta-ref | should | addressed-by-rewrite | these things could be explained in footnotes and/or a respective reference could |
| [108] | 25 | ch2-related-work | beta-ref | should | addressed-by-rewrite | A nice reference could even better back up your claim here. |
| [109] | 26 | ch2-related-work | generic | should | addressed-by-rewrite | Could explain how IaC fits in the overall picture. And how it is combined with C |
| [110] | 26 | ch2-related-work | generic | should | addressed-by-rewrite | Please note that also Terraform also supports this style of configuration/infras |
| [111] | 26 | ch2-related-work | typo | should | addressed-by-rewrite | Thus, the ... |
| [112] | 26 | ch2-related-work | generic | should | addressed-by-rewrite | This is what was also requested in my previous comment ... |
| [113] | 27 | ch2-related-work | generic | should | addressed-by-rewrite | Maybe indicate first what is policy as code and then refer to specific tools tha |
| [114] | 27 | ch2-related-work | generic | should | addressed-by-rewrite | is there any kind of synergy or complementarity wrt PaC and DDT? |
| [115] | 27 | ch2-related-work | generic | should | discussed | reconciliation?adaptation? Please note that I like the second term as it indicat |
| [116] | 27 | ch2-related-work | generic | should | addressed-by-rewrite | ok but in the previous instantiation of the pattern, it is assumed that the onto |
| [117] | 28 | ch2-related-work | generic | should | addressed-by-rewrite | the agent should observe the drift between the actual and desired state and usua |
| [118] | 28 | ch2-related-work | generic | should | addressed-by-rewrite | ok - in this case, there is no reconciliation. Only the gap is detected and reco |
| [119] | 28 | ch2-related-work | generic | should | addressed-by-rewrite | maybe also precision matters here, i.e., that the observations are correct, they |
| [120] | 28 | ch2-related-work | generic | should | addressed-by-rewrite | again these must be correct/precise. |
| [121] | 28 | ch2-related-work | generic | should | addressed-by-rewrite | i.e., the OpenAPI specification's ... |
| [122] | 28 | ch2-related-work | generic | should | addressed-by-rewrite | ok but can they guarantee the correction? Can the correct be deterministic or th |
| [123] | 28 | ch2-related-work | typo | should | addressed-by-rewrite | that |
| [124] | 28 | ch2-related-work | generic | should | addressed-by-rewrite | (which represents the testing report and includes the actionable signals)? |
| [125] | 28 | ch2-related-work | typo | should | addressed-by-rewrite | , for example, ... |
| [126] | 28 | ch2-related-work | generic | should | addressed-by-rewrite | so the focus is on updating the OpenAPI specification, the code or both? The las |
| [127] | 29 | ch2-related-work | generic | should | addressed-by-rewrite | Apart from hiding cloud-native complexity, what are the additional benefits? |
| [128] | 29 | ch2-related-work | generic | should | addressed-by-rewrite | ok but what is the benefit from this structuring? What is the added-value? |
| [129] | 29 | ch2-related-work | generic | should | addressed-by-rewrite | of what? The IDP and the self-service capabilities? |
| [130] | 29 | ch2-related-work | generic | should | addressed-by-rewrite | how and for what purpose? |
| [131] | 29 | ch2-related-work | generic | should | addressed-by-rewrite | what kind of capabilities are meant here? Business capabilities encompassed in s |
| [132] | 29 | ch2-related-work | generic | should | addressed-by-rewrite | ok but a platform consumer is different from the API consumer, right? An API con |
| [133] | 29 | ch2-related-work | generic | should | addressed-by-rewrite | So, 6 principles map directly to the OpenAPI specification and its quality/compl |
| [134] | 29 | ch2-related-work | generic | should | addressed-by-rewrite | correct, if the focus is on the specification. If the focus is on the code, is t |
| [135] | 30 | ch2-related-work | generic | should | addressed-by-rewrite | are these supplied by humans? Are there standards that can be followed for them? |
| [136] | 30 | ch2-related-work | beta-ref | should | addressed-by-rewrite | Is there any reference that formally describes this protocol?  There is a need t |
| [137] | 30 | ch2-related-work | generic | should | addressed-by-rewrite | ok but is this complementary with respect to the respective files that need to b |
| [138] | 31 | ch2-related-work | generic | should | addressed-by-rewrite | again the focus seems to be on the specifications? But the code is equally impor |
| [139] | 31 | ch2-related-work | beta-ref | should | addressed-by-rewrite | could explain how these were derived. A respective reference could make this sel |
| [140] | 31 | ch2-related-work | generic | should | discussed | Ok but there can be an overlap between spec quality assessment and static analys |
| [141] | 31 | ch2-related-work | beta-ref | should | addressed-by-rewrite | please provide a reference that backs up this claim. |
| [142] | 32 | ch2-related-work | generic | should | addressed-by-rewrite | Method. -> mapping to methodology |
| [143] | 32 | ch2-related-work | generic | should | done | Someone can understand what tick and X signify in the table. However, it is not  |
| [144] | 32 | ch2-related-work | generic | should | addressed-by-rewrite | ok but isn't the contract yet another specification? Couldn't the contract also  |
| [145] | 32 | ch2-related-work | generic | should | addressed-by-rewrite | This is the contract-first, API-first or specification-first approach! Most well |
| [146] | 32 | ch2-related-work | generic | should | addressed-by-rewrite | ok - so call it specification-first to be more close to this term! |
| [147] | 32 | ch2-related-work | generic | should | addressed-by-rewrite | So, it is more natural to rely on it, right? Could indicate that in the sentence |
| [148] | 33 | ch2-related-work | generic | should | addressed-by-rewrite | research? |
| [149] | 33 | ch2-related-work | generic | should | addressed-by-rewrite | ok - but were all of these gaps more or less covered in Section 2.6? Or propagat |
| [150] | 33 | ch2-related-work | generic | should | addressed-by-rewrite | complete?informed? Please add the right term/adjective? |
| [151] | 33 | ch2-related-work | generic | should | addressed-by-rewrite | and semantically-rich? |
| [152] | 33 | ch2-related-work | beta-ref | should | addressed-by-rewrite | what is this? Please clarify in footnote |
| [153] | 33 | ch2-related-work | generic | should | addressed-by-rewrite | Ok but based on the analysis, it seems that this gap overlaps with quality asses |
| [154] | 34 | ch3-methodology | beta-name | must | addressed-by-rewrite | why not other principles (from the 9) that target the OpenAPI specification (lik |
| [155] | 34 | ch3-methodology | beta-diag | must | done | Gap 3 -> always use capital for the word (like image, table, section, etc.) |
| [156] | 34 | ch3-methodology | beta-name | should | addressed-by-rewrite | ok but how this relates to another gap? As you establish gap correlations here. |
| [157] | 34 | ch3-methodology | beta-name | should | addressed-by-rewrite | See previous comment ... It seems that you go beyond correlations to justify the |
| [158] | 34 | ch3-methodology | beta-name | should | addressed-by-rewrite | ok but also relates to gap 4? |
| [159] | 35 | ch3-methodology | beta-name | must | done | The Documentation-Driven Methodology  ++ Maybe call it Specification-Driven Test |
| [160] | 35 | ch3-methodology | beta-name | must | done | Actually what you are defining is a conceptual framework for a formal methodolog |
| [161] | 35 | ch3-methodology | beta-name | must | done | ok but I am a little bit puzzled as here you say "complete" while the principles |
| [162] | 35 | ch3-methodology | beta-name | should | addressed-by-rewrite | validation results? |
| [163] | 35 | ch3-methodology | beta-name | must | done | As indicated in a previous comment, I would also add preciseness or accuracy as |
| [164] | 36 | ch3-methodology | beta-name | must | addressed-by-rewrite | in the context of a specific principle? I propose to add this as I presume that  |
| [165] | 36 | ch3-methodology | beta-name | should | deferred | Food for thought: does the OpenAPI cover all possible information? In order to c |
| [166] | 36 | ch3-methodology | beta-name | should | addressed-by-rewrite | I would also add parameter constraints (unless you regard that such constraints  |
| [167] | 36 | ch3-methodology | beta-name | should | addressed-by-rewrite | and is important in the context of API selection & usage/integration? |
| [168] | 36 | ch3-methodology | beta-name | should | addressed-by-rewrite | interpretation? Might be better than "reading". |
| [169] | 37 | ch3-methodology | beta-name | should | addressed-by-rewrite | still this does not check whether the incorporated information is correct, i.e., |
| [170] | 37 | ch3-methodology | beta-name | should | addressed-by-rewrite | do you mean the same API implementation?  In the sense that one specification ca |
| [171] | 37 | ch3-methodology | beta-name | should | addressed-by-rewrite | why isn't non-functional testing also covered? As this depends on the context? |
| [172] | 37 | ch3-methodology | beta-name | must | done | Question: if an API produces different responses/outputs for the same test in di |
| [173] | 37 | ch3-methodology | beta-name | should | addressed-by-rewrite | I agree here. This relates to the concept of flaky tests. Maybe this could be co |
| [174] | 38 | ch3-methodology | beta-name | must | done | as you state axioms as requirements, you could slightly change the wording. E.g. |
| [175] | 38 | ch3-methodology | beta-name | should | addressed-by-rewrite | this might be however verbose - the others might not be (0/1, class or check res |
| [176] | 38 | ch3-methodology | beta-name | must | done | ok but the functional testing does not contribute to observability?  Someone wou |
| [177] | 38 | ch3-methodology | beta-name | must | addressed-by-rewrite | ok but the more kinds of actors are involved, the more are the expectations and  |
| [178] | 38 | ch3-methodology | beta-name | should | addressed-by-rewrite | and analysis? As later on you talk about trends ... |
| [179] | 38 | ch3-methodology | beta-name | should | addressed-by-rewrite | what is meant by "process" here? Please clarify in the text. |
| [180] | 38 | ch3-methodology | beta-name | must | done | ok but before analysing the axioms, you should explain the semantics of the seve |
| [181] | 38 | ch3-methodology | beta-name | should | addressed-by-rewrite | , which have complete theoretical definitions but their implementations are pend |
| [182] | 38 | ch3-methodology | beta-name | should | addressed-by-rewrite | are these standard checks? Are they applied by linters, for instance? You need t |
| [183] | 39 | ch3-methodology | beta-name | should | addressed-by-rewrite | according to the designated version? As it can be the case that a specification  |
| [184] | 39 | ch3-methodology | beta-name | should | addressed-by-rewrite | not clear what is the difference. Can you clarify this? |
| [185] | 39 | ch3-methodology | beta-name | should | addressed-by-rewrite | the explanations aren't the same as the description?  Maybe it is better to indi |
| [186] | 39 | ch3-methodology | beta-name | should | addressed-by-rewrite | examples can also cover parameters, especially when their expected value is not  |
| [187] | 39 | ch3-methodology | beta-name | must | addressed-by-rewrite | where? Maybe you should indicate where this information can be found. This must  |
| [188] | 39 | ch3-methodology | beta-name | should | addressed-by-rewrite | Please again note that this is more about information richness at the structural |
| [189] | 39 | ch3-methodology | beta-name | should | addressed-by-rewrite | and example values? |
| [190] | 39 | ch3-methodology | beta-name | should | addressed-by-rewrite | what does this mean? Please clarify in the text (or at least in a footnote). |
| [191] | 40 | ch3-methodology | beta-name | should | addressed-by-rewrite | who enforces this mode? Can this be a configuration parameter? Please clarify in |
| [192] | 40 | ch3-methodology | beta-name | must | done | there is also a test-ready mode. I do not see sth about it. So, this means that |
| [193] | 40 | ch3-methodology | beta-name | should | addressed-by-rewrite | I agree here. Still what we see is structural. So, it does not convey to the sem |
| [194] | 40 | ch3-methodology | beta-name | must | done | Is this a matter of completeness or heterogeneity? As it relates to the way erro |
| [195] | 40 | ch3-methodology | beta-name | must | addressed-by-rewrite | As some principles differ in terms of their detailed implementation based on the |
| [196] | 40 | ch3-methodology | beta-name | should | addressed-by-rewrite | structural and informative? Structural is needed in order to not stay at a super |
| [197] | 40 | ch3-methodology | beta-name | must | addressed-by-rewrite | As in the case of security levels, the modes need also to be formally defined be |
| [198] | 40 | ch3-methodology | typo | should | addressed-by-rewrite | for them. |
| [199] | 40 | ch3-methodology | beta-name | must | addressed-by-rewrite | Question: I presume that severity means ability to test the API. In this context |
| [200] | 40 | ch3-methodology | beta-name | must | addressed-by-rewrite | ok but for the representation schemas, shouldn't we also check the schema struct |
| [201] | 40 | ch3-methodology | beta-name | should | addressed-by-rewrite | Maybe call this value schema while content schema can be called representation s |
| [202] | 40 | ch3-methodology | beta-name | should | addressed-by-rewrite | shouldn't path parameters be required? |
| [203] | 40 | ch3-methodology | beta-name | should | addressed-by-rewrite | ok but these are not checked in terms of their presence (i.e., a representation  |
| [204] | 41 | ch3-methodology | beta-name | must | addressed-by-rewrite | why this can be considered as minimal? Can you justify this? |
| [205] | 41 | ch3-methodology | beta-name | should | addressed-by-rewrite | This check needs proper justification. Why these two schemes are needed? Are the |
| [206] | 41 | ch3-methodology | beta-name | must | addressed-by-rewrite | Maybe this check should also be done for the previous principle? It can be possi |
| [207] | 41 | ch3-methodology | beta-name | should | addressed-by-rewrite | You provide some implementation details. I am not sure these are relevant at thi |
| [208] | 41 | ch3-methodology | beta-name | must | done | I propose to have this status in all principle descriptions, not just the curren |
| [209] | 41 | ch3-methodology | beta-name | must | addressed-by-rewrite | is this an interface implemented by every principle checker? How each checker is |
| [210] | 42 | ch3-methodology | beta-name | must | addressed-by-rewrite | ok but what is the actual difficulty? As I presume that the framework already ha |
| [211] | 42 | ch3-methodology | beta-name | should | addressed-by-rewrite | ok but how was this established? As it could be regarded as difficult. Maybe thi |
| [212] | 42 | ch3-methodology | beta-name | should | deferred | ok but are these performance expectations imprinted in the OpenAPI spec? If yes, |
| [213] | 42 | ch3-methodology | beta-name | should | deferred | please note that metrics need to be precisely defined or at least named such tha |
| [214] | 42 | ch3-methodology | beta-name | should | addressed-by-rewrite | but are these part of the OpenAPI specification or not? As this was regarded as  |
| [215] | 42 | ch3-methodology | beta-name | must | addressed-by-rewrite | But interestingly, the performance testing is context-aware. So, I am not sure t |
| [216] | 42 | ch3-methodology | beta-name | must | addressed-by-rewrite | does this justify the classification or correlation of the principle with the ob |
| [217] | 42 | ch3-methodology | beta-name | must | addressed-by-rewrite | again, why is this considered minimal? Maybe just the existence of version is mi |
| [218] | 43 | ch3-methodology | beta-name | should | addressed-by-rewrite | NO, it is not the API version. It can mean the version of the API specification. |
| [219] | 43 | ch3-methodology | beta-name | should | addressed-by-rewrite | isn't this implementation detail? |
| [220] | 43 | ch3-methodology | beta-name | should | addressed-by-rewrite | maybe this is incomplete as version could be also a domain-specific word in some |
| [221] | 43 | ch3-methodology | beta-name | must | addressed-by-rewrite | I presume that this is the minimal requirement. The other requirement looks like |
| [222] | 43 | ch3-methodology | beta-name | should | addressed-by-rewrite | It is not clear how these are covered through checks. We understand the semantic |
| [223] | 43 | ch3-methodology | beta-name | should | addressed-by-rewrite | So, this is more related to governance practices, right? The main issue here is  |
| [224] | 43 | ch3-methodology | beta-name | must | addressed-by-rewrite | This overlaps with other principles. I do not understand why it has been specifi |
| [225] | 43 | ch3-methodology | beta-name | must | addressed-by-rewrite | I am not convinced about these checks. They are already covered in other princip |
| [226] | 44 | ch3-methodology | beta-name | must | addressed-by-rewrite | the precondition isn't the completeness of the specification? Thus, to me, this  |
| [227] | 44 | ch3-methodology | beta-name | must | addressed-by-rewrite | if deterministic test execution is not guaranteed without this principle, then w |
| [228] | 44 | ch3-methodology | beta-diag | must | done | I do not see any additional detail in figure that is not covered by the table. U |
| [229] | 44 | ch3-methodology | beta-name | must | addressed-by-rewrite | What about P005. It looks to me that it is also related to completeness. |
| [230] | 44 | ch3-methodology | beta-name | must | done | so they also relate to determinism? I ask this as you correlate them only with c |
| [231] | 45 | ch3-methodology | typo | should | addressed-by-rewrite | This should be critical |
| [232] | 45 | ch3-methodology | beta-name | must | addressed-by-rewrite | Not so clear why P009 is really difficult or non-trivial. The checks seem easy t |
| [233] | 45 | ch3-methodology | beta-name | must | addressed-by-rewrite | Ok but in my opinion, this is a doubly-natured principle in the sense that it re |
| [234] | 45 | ch3-methodology | beta-name | should | addressed-by-rewrite | If an API versioning strategy is not backwards compatible, then we have immediat |
| [235] | 45 | ch3-methodology | beta-name | should | addressed-by-rewrite | Very nice - but the definition of the semantics has to be moved earlier in the S |
| [236] | 46 | ch3-methodology | beta-name | should | addressed-by-rewrite | This could be also moved earlier |
| [237] | 46 | ch3-methodology | beta-name | must | addressed-by-rewrite | What I understand is that all principles are always evaluated. The mode influenc |
| [238] | 46 | ch3-methodology | beta-name | must | addressed-by-rewrite | correct but the validation can be minimal in any principle. It does not have to  |
| [239] | 46 | ch3-methodology | beta-name | must | addressed-by-rewrite | In the description of 3 principles (versioning, security and error handling), yo |
| [240] | 46 | ch3-methodology | beta-name | should | addressed-by-rewrite | Please attempt to make this description fully-justified. Same for following. |
| [241] | 46 | ch3-methodology | beta-name | must | addressed-by-rewrite | but they are minimally based on their descriptions/analysis!  Please check the d |
| [242] | 46 | ch3-methodology | beta-name | must | addressed-by-rewrite | I do not really understand this mode. The other principles have to be satisfied  |
| [243] | 47 | ch4-cli-architecture | generic | should | discussed | are you sure that P002,  P003 and P005 are not relevant? I have the impression t |
| [244] | 47 | ch4-cli-architecture | generic | should | discussed | a release shouldn't be checked against the real tests, functional and performanc |
| [245] | 47 | ch4-cli-architecture | generic | should | discussed | I see mainly minimal -> strict -> test-only as more meaningful stages/phases, wh |
| [246] | 48 | ch4-cli-architecture | generic | should | discussed | The DriveBy System.  I prefer this title. The CLI Architecture indicates that yo |
| [247] | 48 | ch4-cli-architecture | beta-diag | must | addressed-by-rewrite | Maybe you also need to explain the overall development process, such as which we |
| [248] | 48 | ch4-cli-architecture | beta-diag | must | addressed-by-rewrite | The presentation only covers the data flow or also the overall system/tool archi |
| [249] | 48 | ch4-cli-architecture | generic | should | addressed-by-rewrite | system "binary" is an implementation-specific term" |
| [250] | 48 | ch4-cli-architecture | generic | should | addressed-by-rewrite | can we consider these as high-level requirements that drove the design and imple |
| [251] | 48 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Ok but there are structural differences between the different versions while new |
| [252] | 48 | ch4-cli-architecture | beta-ref | should | addressed-by-rewrite | you need to provide the design goals in an architecture-agnostic manner. So, ref |
| [253] | 49 | ch4-cli-architecture | generic | should | addressed-by-rewrite | based on the considered principles, this is an almost valid design goal. The sol |
| [254] | 49 | ch4-cli-architecture | generic | should | addressed-by-rewrite | apart from the structuring, is there any other kind of related need?  For instan |
| [255] | 49 | ch4-cli-architecture | generic | should | addressed-by-rewrite | in principle - the content may vary. |
| [256] | 49 | ch4-cli-architecture | generic | should | addressed-by-rewrite | only these two were applied? Any other? |
| [257] | 49 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok - so there is a common interface for checkers that is adopted independently o |
| [258] | 49 | ch4-cli-architecture | beta-flow | must | addressed-by-rewrite | overflow here ... |
| [259] | 49 | ch4-cli-architecture | generic | should | addressed-by-rewrite | you forgot to say sth here about two other packages: testing and loader. |
| [260] | 50 | ch4-cli-architecture | generic | should | addressed-by-rewrite | could also indicate some words about each layer. For instance, what is generally |
| [261] | 50 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Question here: in terms of dependency, it seems that the core tests (functional  |
| [262] | 50 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Nine checkers you have designed, right? Not 8 ones. |
| [263] | 51 | ch4-cli-architecture | generic | should | addressed-by-rewrite | The last three packages were not mentioned before. Can we consider them as ortho |
| [264] | 51 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Suddenly, you provide here some code without explaining what was the implementat |
| [265] | 51 | ch4-cli-architecture | generic | should | deferred | very minor: you have a comment extending (visually) one like while it is a one-l |
| [266] | 52 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Question here: testmode was not introduced in the previous chapter. What is its  |
| [267] | 53 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Ok, here you indicate the overall validation result for a specific principle. Wh |
| [268] | 53 | ch4-cli-architecture | generic | should | addressed-by-rewrite | in general? Or in the context of specific tests? It could be possible to have ca |
| [269] | 53 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Ok - maybe refer to the Appendix to see the definition of this struct? |
| [270] | 53 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok but the checks are not well mapped to the principle results in my opinion. Sh |
| [271] | 54 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Correct - I totally agree here! |
| [272] | 54 | ch4-cli-architecture | beta-flow | must | addressed-by-rewrite | overflow |
| [273] | 54 | ch4-cli-architecture | generic | should | addressed-by-rewrite | very good. If a new feature is involved in the OpenAPI standard that requires th |
| [274] | 54 | ch4-cli-architecture | beta-flow | must | addressed-by-rewrite | Table overflows significantly. Please correct - the third column content is cut  |
| [275] | 54 | ch4-cli-architecture | beta-ref | should | addressed-by-rewrite | please explain in footnote what this means. |
| [276] | 54 | ch4-cli-architecture | generic | should | addressed-by-rewrite | how? Can you provide a respective example? |
| [277] | 55 | ch4-cli-architecture | generic | should | addressed-by-rewrite | alternatively as one of them will be used in the context of a specific OpenAPI s |
| [278] | 55 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Could explain in text at high-level the main logic of this code. |
| [279] | 56 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok - but maybe there could be another way to solve this problem? As sanity check |
| [280] | 56 | ch4-cli-architecture | generic | should | addressed-by-rewrite | what is spec.go? Is the implementation of the ApiSpec? |
| [281] | 57 | ch4-cli-architecture | beta-flow | must | addressed-by-rewrite | overflow again here |
| [282] | 58 | ch4-cli-architecture | beta-flow | must | addressed-by-rewrite | overflow again here. |
| [283] | 58 | ch4-cli-architecture | generic | should | done | so they downgrade into 3.0.x some parts of the document in order to enable its v |
| [284] | 58 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Maybe the justification for adopting this library could be mentioned beforehand  |
| [285] | 58 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Question: a checker implements all tests as presented in previous chapter? I pre |
| [286] | 58 | ch4-cli-architecture | beta-ref | should | addressed-by-rewrite | what does this mean? Maybe explain in footnote. Please note that I have seen con |
| [287] | 59 | ch4-cli-architecture | generic | should | addressed-by-rewrite | question: this method covers both version cases for OpenAPI? So, it also conduct |
| [288] | 60 | ch4-cli-architecture | beta-flow | must | addressed-by-rewrite | Overflow also here |
| [289] | 60 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Please "check" whether the number of checks is correct per principle. |
| [290] | 60 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok - could also justify why P003 is the second largest checker.  And why P008 is |
| [291] | 60 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Here you provide only the implemented ones, right? |
| [292] | 61 | ch4-cli-architecture | typo | should | done | Not implemented yet ... |
| [293] | 61 | ch4-cli-architecture | generic | should | addressed-by-rewrite | what is meant by this? What is a combined mode? Do you mean that someone could s |
| [294] | 61 | ch4-cli-architecture | generic | should | addressed-by-rewrite | so, it is a kind of backend? That interfaces with the CLI (front-end)? |
| [295] | 62 | ch4-cli-architecture | generic | should | addressed-by-rewrite | do we see it somewhere? Could be shown in the Appendix. |
| [296] | 63 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok but this has to be proven. Maybe point to the evaluation chapter here for thi |
| [297] | 63 | ch4-cli-architecture | generic | should | discussed | please note that this covers single interactions. However, in some cases, compos |
| [298] | 63 | ch4-cli-architecture | generic | should | addressed-by-rewrite | how does it find if an operation is deprecated or not? |
| [299] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok but this looks like fallback, it might fail to provide a correct test. |
| [300] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | How are these supplied? Please clarify in the text. |
| [301] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok but this looks like blind validation in the sense that it is not clear what s |
| [302] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | So, skipping indicates synchronisation issues between spec & implementation in t |
| [303] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | "reporting ...": the API reports the requirement in its responses? Please clarif |
| [304] | 64 | ch4-cli-architecture | generic | should | discussed | so, maybe the other testing approaches are not deterministic? That is why you im |
| [305] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok but this is one kind of performance testing and many others exist. I presume  |
| [306] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | What does this mean? Please clarify and also indicate how this is achieved. |
| [307] | 64 | ch4-cli-architecture | generic | should | done | I do not like the term "attack" here. You do not conduct penetration/security te |
| [308] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Ok but how are these defined? Are they part of the spec? Or they are part of the |
| [309] | 64 | ch4-cli-architecture | generic | should | addressed-by-rewrite | Ok but this was already mentioned above for both of these testing kinds. No need |
| [310] | 65 | ch4-cli-architecture | generic | should | addressed-by-rewrite | So, this is "static" information coming from configuration files or command flag |
| [311] | 65 | ch4-cli-architecture | generic | should | addressed-by-rewrite | question: the version or deployment is not covered in the name of the report. So |
| [312] | 65 | ch4-cli-architecture | generic | should | addressed-by-rewrite | So, this is not done by the JSON renderer? Or this is done by both of them? |
| [313] | 65 | ch4-cli-architecture | typo | should | addressed-by-rewrite | test-kind-specific ... |
| [314] | 65 | ch4-cli-architecture | generic | should | discussed | Ok - but maybe a schema could be good to have here? As it gives predictability:  |
| [315] | 65 | ch4-cli-architecture | generic | should | discussed | maybe for now. But in the future, it might also incorporate them - e.g., they co |
| [316] | 65 | ch4-cli-architecture | generic | should | addressed-by-rewrite | ok - so, we always talk about one remediation. Thus, there are not multiple alte |
| [317] | 66 | ch4-cli-architecture | beta-flow | must | addressed-by-rewrite | two overflows exist in this page |
| [318] | 66 | ch4-cli-architecture | beta-diag | must | addressed-by-rewrite | ok - but maybe also indicate what is the file for implementing the CLI itself? M |
| [319] | 67 | ch4-cli-architecture | generic | should | addressed-by-rewrite | if the spec is missing, is this an infrastructure problem? Even if you retry, th |
| [320] | 67 | ch4-cli-architecture | beta-diag | must | addressed-by-rewrite | ok but we did not see architecture and data flow diagrams. Only a description of |
| [321] | 68 | ch5-kubernetes | beta-ref | should | addressed-by-rewrite | question: this verification requires the API to be somewhere available. Should w |
| [322] | 69 | ch5-kubernetes | generic | should | addressed-by-rewrite | Please remove "of Chapter 2". It is obvious. |
| [323] | 70 | ch5-kubernetes | generic | should | addressed-by-rewrite | ok but how is this achieved?Do you have different binary versions per each OS? O |
| [324] | 70 | ch5-kubernetes | generic | should | addressed-by-rewrite | are these two involved in the next stage and the last one is about the final sta |
| [325] | 71 | ch5-kubernetes | typo | should | done | he/she declares |
| [326] | 71 | ch5-kubernetes | typo | should | addressed-by-rewrite | no matter how ... |
| [327] | 71 | ch5-kubernetes | typo | should | addressed-by-rewrite | a |
| [328] | 72 | ch5-kubernetes | beta-diag | must | addressed-by-rewrite | It is not clear what are these abstraction layers. Can you depict a nice figure  |
| [329] | 72 | ch5-kubernetes | generic | should | addressed-by-rewrite | where is this layer placed? It is the highest possible? |
| [330] | 72 | ch5-kubernetes | beta-flow | must | addressed-by-rewrite | Again, this table overflows. Please fix this |
| [331] | 72 | ch5-kubernetes | beta-diag | must | addressed-by-rewrite | was there something implementation-specific that must be highlighted here? Or th |
| [332] | 72 | ch5-kubernetes | generic | should | addressed-by-rewrite | ok but from where do these come from? You need to clarify how you derive your go |
| [333] | 73 | ch5-kubernetes | beta-ref | should | addressed-by-rewrite | how these correlate with the validation logic that was implemented in the CLI? I |
| [334] | 73 | ch5-kubernetes | generic | should | addressed-by-rewrite | can you please elaborate more on this? Which is also higher than the other? |
| [335] | 73 | ch5-kubernetes | typo | should | addressed-by-rewrite | (developed by the platform team) |
| [336] | 73 | ch5-kubernetes | beta-diag | must | done | What does system context mean? Is this a kind of context diagram that you are pr |
| [337] | 74 | ch5-kubernetes | beta-diag | must | done | It is also not clear in this figure what is already there and what was implement |
| [338] | 74 | ch5-kubernetes | beta-diag | must | done | do the colours play a role in the diagram? If yes, please indicate their semanti |
| [339] | 74 | ch5-kubernetes | generic | should | addressed-by-rewrite | ok, I agree. But the key is the state. It has to be permanent and transactional, |
| [340] | 75 | ch5-kubernetes | generic | should | addressed-by-rewrite | is this a component of a Kubernetes controller or the controller itself? As you  |
| [341] | 75 | ch5-kubernetes | typo | should | done | gradually |
| [342] | 75 | ch5-kubernetes | typo | should | addressed-by-rewrite | an |
| [343] | 76 | ch5-kubernetes | beta-ref | should | addressed-by-rewrite | Ok but these are provider-specific for the Github provider. So, here you supply  |
| [344] | 76 | ch5-kubernetes | beta-flow | must | addressed-by-rewrite | again overflow here |
| [345] | 76 | ch5-kubernetes | beta-ref | should | addressed-by-rewrite | what does this mean? Please clarify in footnote |
| [346] | 77 | ch5-kubernetes | generic | should | addressed-by-rewrite | Is this the actual implementation that you provide to realise all the magic? If  |
| [347] | 77 | ch5-kubernetes | generic | should | addressed-by-rewrite | here the arrows showcase dependencies and order of creation or sth else? |
| [348] | 77 | ch5-kubernetes | generic | should | addressed-by-rewrite | earlier? When was that attempted and when made you change your mind? |
| [349] | 78 | ch5-kubernetes | generic | should | addressed-by-rewrite | What is this? It is not clear to the reader. |
| [350] | 78 | ch5-kubernetes | typo | should | done | gitops vs GitOps |
| [351] | 78 | ch5-kubernetes | generic | should | addressed-by-rewrite | what is the semantics of these branches? Please clarify |
| [352] | 78 | ch5-kubernetes | generic | should | addressed-by-rewrite | again this is not clear - very specific terminology that might not be clear to t |
| [353] | 78 | ch5-kubernetes | beta-diag | must | addressed-by-rewrite | image pull of what? |
| [354] | 79 | ch5-kubernetes | generic | should | addressed-by-rewrite | ok but what is a check? Does it relate to the principle checks that we show in t |
| [355] | 79 | ch5-kubernetes | generic | should | addressed-by-rewrite | there is also the test-only check available in the CLI that is a combination of  |
| [356] | 79 | ch5-kubernetes | beta-flow | must | addressed-by-rewrite | Overflow issue here |
| [357] | 79 | ch5-kubernetes | generic | should | addressed-by-rewrite | which subset? |
| [358] | 79 | ch5-kubernetes | beta-diag | must | addressed-by-rewrite | which configures the validation? So, provides values for respective configuratio |
| [359] | 80 | ch5-kubernetes | beta-flow | must | addressed-by-rewrite | overflow + now it is clear what is rate. But this load-test maps to a single loa |
| [360] | 80 | ch5-kubernetes | generic | should | addressed-by-rewrite | ok but the ordering is required by the model or it is enforced by the implementa |
| [361] | 80 | ch5-kubernetes | generic | should | addressed-by-rewrite | So there are applied by default? Can't the user modify this? As this should be c |
| [362] | 80 | ch5-kubernetes | generic | should | addressed-by-rewrite | why aren't functional-tests also executed for the production gate? Should we con |
| [363] | 81 | ch5-kubernetes | generic | should | addressed-by-rewrite | so these two are fixed while the rest of the information is configurable by the  |
| [364] | 81 | ch5-kubernetes | generic | should | addressed-by-rewrite | ok - but we should also expect that configurations can be given also by the CRD? |
| [365] | 82 | ch5-kubernetes | generic | should | addressed-by-rewrite | Are we missing per-environment resources here? |
| [366] | 83 | ch5-kubernetes | typo | should | addressed-by-rewrite | , respectively, |
| [367] | 83 | ch5-kubernetes | generic | should | addressed-by-rewrite | which helm chart? Is there Helm chart in your repo? Please clarify. |
| [368] | 83 | ch5-kubernetes | generic | should | addressed-by-rewrite | is there any assumption about this installation? What should be already in place |
| [369] | 83 | ch5-kubernetes | beta-diag | must | addressed-by-rewrite | please also provide reference to the figure that shows the CRD model. |
| [370] | 83 | ch5-kubernetes | generic | should | done | please validate whether it is 44 or 47. I have the impression that it can be 47 |
| [371] | 83 | ch5-kubernetes | beta-flow | must | addressed-by-rewrite | overflow here |
| [372] | 83 | ch5-kubernetes | generic | should | addressed-by-rewrite | ok - so it does the monitoring. The reconciliation is done by which function? Th |
| [373] | 84 | ch5-kubernetes | generic | should | addressed-by-rewrite | the arrow here has the right direction? As the function should observe the resou |
| [374] | 84 | ch5-kubernetes | generic | should | addressed-by-rewrite | by mistake or on purpose? |
| [375] | 84 | ch5-kubernetes | generic | should | addressed-by-rewrite | can this be changed by the user? Or these are hardcoded? |
| [376] | 85 | ch5-kubernetes | typo | should | addressed-by-rewrite | a |
| [377] | 87 | ch5-kubernetes | generic | should | addressed-by-rewrite | (see Section 5.6.2 below) |
| [378] | 87 | ch5-kubernetes | generic | should | addressed-by-rewrite | if there is a change, does it inform the previous component? Shouldn't that comp |
| [379] | 87 | ch5-kubernetes | generic | should | addressed-by-rewrite | function-go? |
| [380] | 87 | ch5-kubernetes | beta-flow | must | addressed-by-rewrite | overflow + maybe this is moved to the description of the function? |
| [381] | 88 | ch5-kubernetes | beta-flow | must | addressed-by-rewrite | please be aware of the overflows which happen in almost every page |
| [382] | 89 | ch5-kubernetes | generic | should | addressed-by-rewrite | so, this is a kind of configuration point? |
| [383] | 89 | ch5-kubernetes | generic | should | addressed-by-rewrite | Maybe in Chapter 2 could provide more explanations about them in the context of  |
| [384] | 89 | ch5-kubernetes | generic | should | addressed-by-rewrite | so, these are the global configurations/settings ... |
| [385] | 90 | ch5-kubernetes | generic | should | addressed-by-rewrite | default or the one that has been applied during the installation (as defaults ca |
| [386] | 90 | ch5-kubernetes | generic | should | addressed-by-rewrite | this was not mentioned before. So, it is also not clear what was involved in tha |
| [387] | 91 | ch5-kubernetes | generic | should | deferred | could this be an extension of your work to satisfy different potential "business |
| [388] | 93 | ch5-kubernetes | typo | should | addressed-by-rewrite | mapping to |
| [389] | 93 | ch5-kubernetes | generic | should | addressed-by-rewrite | this is the name of the step? I am asking as the name is given in an example tha |
| [390] | 93 | ch5-kubernetes | beta-flow | must | addressed-by-rewrite | overflow again here |
| [391] | 93 | ch5-kubernetes | typo | should | done | (PersistentVolumeClaim) |
| [392] | 93 | ch5-kubernetes | generic | should | addressed-by-rewrite | this runs where? |
| [393] | 94 | ch5-kubernetes | generic | should | addressed-by-rewrite | against the live API? |
| [394] | 94 | ch5-kubernetes | generic | should | addressed-by-rewrite | chart installation? |
| [395] | 95 | ch5-kubernetes | generic | should | addressed-by-rewrite | Maybe clarify that here you start indicating how the workflow template is dynami |
| [396] | 95 | ch5-kubernetes | typo | should | addressed-by-rewrite | For example, ... |
| [397] | 96 | ch5-kubernetes | beta-diag | must | addressed-by-rewrite | so these will be some checks depending on the validation mode (default or config |
| [398] | 96 | ch5-kubernetes | generic | should | addressed-by-rewrite | ok but the thresholds apply mainly for the load-test check. So, this holds only  |
| [399] | 96 | ch5-kubernetes | generic | should | addressed-by-rewrite | did not explain how this was produced and what is its content. |
| [400] | 97 | ch5-kubernetes | generic | should | addressed-by-rewrite | and applied? |
| [401] | 97 | ch5-kubernetes | generic | should | addressed-by-rewrite | I presume that the installation covers the first three steps.  The application o |
| [402] | 97 | ch5-kubernetes | generic | should | addressed-by-rewrite | Yes but this is not installation step per se and should not be advertised as suc |
| [403] | 98 | ch6-gitops | generic | should | addressed-by-rewrite | Nothing is mentioned about Sections 5.2. |
| [404] | 99 | ch6-gitops | beta-ref | should | addressed-by-rewrite | if each element maps to a section, could reference that section in parenthesis.  |
| [405] | 99 | ch6-gitops | generic | should | addressed-by-rewrite | in Chapter X ... Or is this done in this chapter? |
| [406] | 100 | ch6-gitops | typo | should | addressed-by-rewrite | The aforementioned |
| [407] | 100 | ch6-gitops | generic | should | addressed-by-rewrite | which is guaranteed? Please clarify |
| [408] | 100 | ch6-gitops | generic | should | discussed | Maybe the vertical labels should be moved to the middle of each rectangle. This  |
| [409] | 100 | ch6-gitops | beta-diag | must | addressed-by-rewrite | A has been "consumed" in the figure. Please fix this (Argo Events) |
| [410] | 100 | ch6-gitops | generic | should | addressed-by-rewrite | The arrow here is correct? Same for arrow connecting Sensor and Argo Workflow. |
| [411] | 100 | ch6-gitops | generic | should | addressed-by-rewrite | The rectangle length should be reduced so as to be within the overall boundary.  |
| [412] | 101 | ch6-gitops | typo | should | addressed-by-rewrite | As the ... |
| [413] | 101 | ch6-gitops | beta-ref | should | addressed-by-rewrite | and by knowing the aforementioned URL pattern |
| [414] | 101 | ch6-gitops | beta-flow | must | addressed-by-rewrite | overflow ... |
| [415] | 102 | ch6-gitops | generic | should | discussed | If you refer to Section 7.3 that more or less conducts the latency analysis, the |
| [416] | 102 | ch6-gitops | generic | should | addressed-by-rewrite | ok but in principle, the test time can be greater depending as you indicate on s |
| [417] | 103 | ch6-gitops | beta-ref | should | addressed-by-rewrite | Is there a reference that also clearly indicates this - that it is a bad practic |
| [418] | 103 | ch6-gitops | beta-flow | must | addressed-by-rewrite | Lots of overflows in the current page |
| [419] | 104 | ch6-gitops | typo | should | addressed-by-rewrite | As each ... |
| [420] | 105 | ch6-gitops | generic | should | addressed-by-rewrite | This is a very big advantage. I would "advertise" it as such in order to highlig |
| [421] | 105 | ch6-gitops | generic | should | addressed-by-rewrite | I would recommend to better explain how hydration works in Chapter 2 (what is co |
| [422] | 106 | ch6-gitops | beta-diag | must | addressed-by-rewrite | Figure is nice. Two improvements: (a) "manifests ..." label could be placed cent |
| [423] | 107 | ch6-gitops | typo | should | done | the labels here overlap! |
| [424] | 107 | ch6-gitops | generic | should | addressed-by-rewrite | The arrows in this rectangle and the "DDT ..." one are misleading. Logically spe |
| [425] | 107 | ch6-gitops | beta-flow | must | addressed-by-rewrite | again overflow here. Please note that due to the overflow in many pages the info |
| [426] | 107 | ch6-gitops | generic | should | addressed-by-rewrite | must be formatted? What does it mean that it must match these two things? |
| [427] | 108 | ch6-gitops | generic | should | addressed-by-rewrite | ok but you need to stress that the order of the environments is important as it  |
| [428] | 108 | ch6-gitops | generic | should | addressed-by-rewrite | which workflow? |
| [429] | 108 | ch6-gitops | beta-ref | should | addressed-by-rewrite | again this is nice. The vision could be to support other providers apart from Gi |
| [430] | 109 | ch6-gitops | beta-diag | must | addressed-by-rewrite | this could be also shown with a nice figure, if possible |
| [431] | 109 | ch6-gitops | generic | should | addressed-by-rewrite | very nice section - this is what I would like to see as a reader and as a potent |
| [432] | 110 | ch6-gitops | generic | should | addressed-by-rewrite | which signifies automated sync. |
| [433] | 110 | ch6-gitops | typo | should | addressed-by-rewrite | ; it signifies ... |
| [434] | 110 | ch6-gitops | generic | should | addressed-by-rewrite | ; this signifies ... along with the existence of a gate.. |
| [435] | 110 | ch6-gitops | generic | should | addressed-by-rewrite | but this can be also applied. There is no quality gate specified and autoMerge:  |
| [436] | 110 | ch6-gitops | generic | should | addressed-by-rewrite | ok but does it make sense? Which cases could it cover? Maybe it is not recommend |
| [437] | 111 | ch6-gitops | generic | should | addressed-by-rewrite | Again some arrows do not have straight lines |
| [438] | 111 | ch6-gitops | beta-diag | must | addressed-by-rewrite | ok but one configuration (no gate, no auto-merge) is not covered by the figure.  |
| [439] | 112 | ch6-gitops | generic | should | addressed-by-rewrite | functional testing is not covered in the sentence |
| [440] | 112 | ch6-gitops | generic | should | addressed-by-rewrite | usually the Setpoint is compared against a variable. What is the variable in you |
| [441] | 112 | ch6-gitops | generic | should | discussed | I do not have a problem with OpenAPI specification being the Setpoint.  However, |
| [442] | 112 | ch6-gitops | generic | should | addressed-by-rewrite | (approve) -> he/she should manually approve the promotions. This is a major deve |
| [443] | 112 | ch6-gitops | generic | should | addressed-by-rewrite | The feedback in control theory signifies if the current situation is ok or not.  |
| [444] | 113 | ch6-gitops | beta-diag | must | addressed-by-rewrite | Nice figure but there are issues as labels are not shown well and arrows are hid |
| [445] | 113 | ch6-gitops | generic | should | addressed-by-rewrite | correct - but based on my previous comments, the specification must be complete  |
| [446] | 113 | ch6-gitops | generic | should | discussed | ok but the issue is that the workflow has failed. Shouldn't it be re-run in orde |
| [447] | 113 | ch6-gitops | generic | should | discussed | these are not the three well-known control-theory properties: stability, control |
| [448] | 113 | ch6-gitops | generic | should | addressed-by-rewrite | ok but the question is whether this holds for the setpoint. The checks that you  |
| [449] | 114 | ch6-gitops | generic | should | addressed-by-rewrite | ok but if document(ation) is the pointset, it has to be complete. But I agree th |
| [450] | 114 | ch6-gitops | generic | should | addressed-by-rewrite | not clear what you are trying to identify here. Sth mentioned in the documentati |
| [451] | 115 | ch6-gitops | generic | should | addressed-by-rewrite | main implementation logic (as now we talk about source code and its CI/CD) |
| [452] | 116 | ch6-gitops | beta-ref | should | done | what is KubeCore? I am not sure this was indicated before. Maybe dedicate a foot |
| [453] | 117 | ch6-gitops | generic | should | addressed-by-rewrite | (so only all relevant resources) |
| [454] | 119 | ch7-evaluation | gamma | must | addressed-by-rewrite | Did not say sth about Section 6.8. |
| [455] | 120 | ch7-evaluation | gamma | must | done | If PoC evaluation does not lead to assessing various metrics, I would recommend |
| [456] | 120 | ch7-evaluation | gamma | must | addressed-by-rewrite | Ok but aren't the evaluations also showing this? That an ontological specificati |
| [457] | 121 | ch7-evaluation | gamma | must | addressed-by-rewrite | Question: do we see the PoC operational evaluation? Is this this last arm? I am  |
| [458] | 121 | ch7-evaluation | gamma | must | addressed-by-rewrite | detection accuracy is measured by which metrics? Before presenting the controlle |
| [459] | 121 | ch7-evaluation | gamma | must | addressed-by-rewrite | complete -> in the sense of specification completeness, which is necessary to gu |
| [460] | 121 | ch7-evaluation | gamma | must | addressed-by-rewrite | what is meant by a task? |
| [461] | 121 | ch7-evaluation | gamma | must | addressed-by-rewrite | are you sure about that? I believe that the large-scale evaluation does not rely |
| [462] | 122 | ch7-evaluation | gamma | must | addressed-by-rewrite | could explain how this is done.Do you have a different branch per defect injecti |
| [463] | 123 | ch7-evaluation | gamma | must | done | is this expected? I mean that you indicate here what is the expectation by evalu |
| [464] | 123 | ch7-evaluation | gamma | must | addressed-by-rewrite | ok but what is really meant by quality dimension? As we know that there are nice |
| [465] | 123 | ch7-evaluation | gamma | must | done | Very interesting aspect to include OpenAPI version in the table. However, you se |
| [466] | 123 | ch7-evaluation | gamma | must | done | If you have removed a specific quality dimension from here, the naming of the AP |
| [467] | 123 | ch7-evaluation | gamma | must | done | Here the principle affected seems to be P005. The same holds for the next API. B |
| [468] | 123 | ch7-evaluation | gamma | must | addressed-by-rewrite | ok but it is not clear how degradation is detected in this case. You put a speci |
| [469] | 123 | ch7-evaluation | gamma | must | addressed-by-rewrite | ok - are these two also critical or not? Because you state in the text below tha |
| [470] | 123 | ch7-evaluation | gamma | must | addressed-by-rewrite | but in the table, maybe all defects seem to be blocking? |
| [471] | 124 | ch7-evaluation | gamma | must | done | I disagree here with the table content in the sense that it has to cover everyth |
| [472] | 124 | ch7-evaluation | gamma | must | done | I am puzzled! The perfect-api as was presented in previous section does not exhi |
| [473] | 124 | ch7-evaluation | gamma | must | addressed-by-rewrite | In the previous table, you indicated that bad-docs-api suffers from a security-o |
| [474] | 124 | ch7-evaluation | gamma | must | addressed-by-rewrite | but as I understand, this is not the perfect-api that was mentioned in the previ |
| [475] | 124 | ch7-evaluation | gamma | must | done | By reading this paragraph, I am more puzzled than before about what is the perfe |
| [476] | 124 | ch7-evaluation | gamma | must | addressed-by-rewrite | I would indicate that non-critical-api has been simulated to match current devel |
| [477] | 124 | ch7-evaluation | gamma | must | addressed-by-rewrite | no-auth-api was said to be missing authentication scheme and bad-docs-api a secu |
| [478] | 124 | ch7-evaluation | gamma | must | addressed-by-rewrite | so these were the two faults injected in the broken-api? I believe that these sh |
| [479] | 125 | ch7-evaluation | gamma | must | addressed-by-rewrite | yes but this must have been already communicated. |
| [480] | 126 | ch7-evaluation | gamma | must | addressed-by-rewrite | I propose to have 4 layers so that you have another evidence that the gate mecha |
| [481] | 126 | ch7-evaluation | gamma | must | discussed | 4/9 as there are 4 principles that are critical |
| [482] | 126 | ch7-evaluation | gamma | must | addressed-by-rewrite | I believe that as you mention slow-api that concerns the performance testing, si |
| [483] | 126 | ch7-evaluation | gamma | must | addressed-by-rewrite | ok but still the issues that exhibits can influence functional testing. Thus, so |
| [484] | 127 | ch7-evaluation | gamma | must | addressed-by-rewrite | In this respect, the idea is that there are two cases that can enable promotion  |
| [485] | 127 | ch7-evaluation | gamma | must | done | Need to provide precise number of APIs downloaded. |
| [486] | 128 | ch7-evaluation | gamma | must | addressed-by-rewrite | The OAUth2 scope is needed when type: oauth2 (or OpenID Connect built on top of  |
| [487] | 128 | ch7-evaluation | gamma | must | done | I do not like this title here for the column. As you have done the evaluation, i |
| [488] | 128 | ch7-evaluation | gamma | must | done | What does high, low and medium mean in practice? High is >= 90%, medium is >= 50 |
| [489] | 128 | ch7-evaluation | gamma | must | addressed-by-rewrite | It would be nice to move the focus also on each principle in order to see which  |
| [490] | 128 | ch7-evaluation | gamma | must | done | Did not give any comment about P005 and where the pass rate could be considered |
| [491] | 129 | ch7-evaluation | gamma | must | addressed-by-rewrite | partially in my opinion: it matches P001 but matching fails for P002-P004 in the |
| [492] | 129 | ch7-evaluation | gamma | must | addressed-by-rewrite | were not passed by any of the APIs |
| [493] | 129 | ch7-evaluation | gamma | must | addressed-by-rewrite | to me 80% is high rather than intermediate. |
| [494] | 129 | ch7-evaluation | gamma | must | done | I believe that there is a need to not provide similar information as already sup |
| [495] | 129 | ch7-evaluation | gamma | must | done | Please explain the structure of this section in one paragraph and correlate it w |
| [496] | 129 | ch7-evaluation | gamma | must | addressed-by-rewrite | which one? |
| [497] | 130 | ch7-evaluation | gamma | must | addressed-by-rewrite | ok but of course the transition from one environment to the other requires manua |
| [498] | 130 | ch7-evaluation | gamma | must | addressed-by-rewrite | What do you mean by "declared specification"? I believe that what you are trying |
| [499] | 130 | ch7-evaluation | gamma | must | addressed-by-rewrite | This relates to P006 and P007 while here you talk about the other principles tha |
| [500] | 131 | ch7-evaluation | gamma | must | addressed-by-rewrite | API testing isn't a form of integration testing? So, here you mean "other kinds  |
| [501] | 131 | ch7-evaluation | gamma | must | addressed-by-rewrite | Indeed, but in this kind of testing, not just functional (API) testing but many  |
| [502] | 131 | ch7-evaluation | gamma | must | addressed-by-rewrite | sure but instead of load testing, maybe it is better to do stress testing? |
| [503] | 131 | ch7-evaluation | gamma | must | addressed-by-rewrite | Ok but may this leads to the requirement to have two control loops instead of on |
| [504] | 132 | ch7-evaluation | gamma | must | addressed-by-rewrite | Not clear how these were derived. Need to explain the respective rationale. Mayb |
| [505] | 132 | ch7-evaluation | gamma | must | addressed-by-rewrite | Ok but what does it mean pass and fail?  Pass means that the problem passed the  |
| [506] | 132 | ch7-evaluation | gamma | must | addressed-by-rewrite | Is this really an issue? As the system has been designed deliberately to address |
| [507] | 132 | ch7-evaluation | gamma | must | addressed-by-rewrite | I presume that this scenario shows that the system can scale to cover multiple r |
| [508] | 132 | ch7-evaluation | gamma | must | addressed-by-rewrite | you mean the ones used in the first experiment/validation? Please clarify |
| [509] | 132 | ch7-evaluation | gamma | must | done | I assume that this covers a complete history / period of real usage of the repos |
| [510] | 132 | ch7-evaluation | gamma | must | done | These projects/repo were never introduced before. I believe that there is a need |
| [511] | 132 | ch7-evaluation | gamma | must | addressed-by-rewrite | So, this is sth that needs to be corrected in the future? How this could be done |
| [512] | 133 | ch7-evaluation | gamma | must | addressed-by-rewrite | Please also indicate how these metrics were computed. If multiple tools were uti |
| [513] | 133 | ch7-evaluation | gamma | must | addressed-by-rewrite | you mean of the tool when run in standalone manner? Maybe you can add a column t |
| [514] | 133 | ch7-evaluation | gamma | must | addressed-by-rewrite | This could be inspected by making observations during the large-scale experiment |
| [515] | 133 | ch7-evaluation | gamma | must | addressed-by-rewrite | does this depend also on the computing power of the cluster? Or can we consider  |
| [516] | 133 | ch7-evaluation | gamma | must | addressed-by-rewrite | Could indicate here the factors that can influence this metric as done for valid |
| [517] | 133 | ch7-evaluation | gamma | must | addressed-by-rewrite | could you have respective estimates for manual configuration? This could make a  |
| [518] | 133 | ch7-evaluation | gamma | must | addressed-by-rewrite | (Related to RQ1) Please put this in each relevant parenthesis as otherwise you m |
| [519] | 133 | ch7-evaluation | gamma | must | addressed-by-rewrite | Seems that paragraph is not fully justified +  Of course, the experiment could b |
| [520] | 134 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | Correct API addressing & classification? I propose just a new title to strengthe |
| [521] | 134 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | Correlation to RQ2 is not quite evident and justified in the paragraph. Please t |
| [522] | 134 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | it depends on what is inside the layer. If the layer covers everything, then it  |
| [523] | 134 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | which are the 4 defect types? Invalid API (P001), critical quality issues (P005) |
| [524] | 134 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | Maybe change to: "Automation extent confirmed due to validation approach general |
| [525] | 134 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | I do not understand how quality gaps connect to RQ2 which is about automation ex |
| [526] | 134 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | ok, what does specification mean? |
| [527] | 134 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | RQ3 + RQ4 as RQ4 is also proven/validated based on PoC as the ontological specif |
| [528] | 135 | ch8-ai-assisted | alpha | must | done | Please update the chapter to better clarify the own effort that you have put in |
| [529] | 135 | ch8-ai-assisted | typo | should | addressed-by-rewrite | software |
| [530] | 136 | ch8-ai-assisted | typo | should | addressed-by-rewrite | ? |
| [531] | 136 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | why do you say "later" here? What was the initial situation? |
| [532] | 137 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | Ok but maybe you could explain why did you follow this structure and how many le |
| [533] | 137 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | ok but each file implements a specific principle, right? Please clarify. |
| [534] | 137 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | I do not understand this as previously you have indicated that an agent acting o |
| [535] | 138 | ch8-ai-assisted | alpha | must | done | I would indicate below in text which are these code quality rules. |
| [536] | 138 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | where do you indicate which principles are critical or not? How and where valida |
| [537] | 138 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | so, by API inventory you mean the repo with APIs that you have collected from AP |
| [538] | 138 | ch8-ai-assisted | beta-diag | must | addressed-by-rewrite | this includes figures you would like to have apart from those to be generated by |
| [539] | 138 | ch8-ai-assisted | beta-diag | must | addressed-by-rewrite | Not clear what is to be included in the documentation. Please clarify. The diagr |
| [540] | 138 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | CSV is used for what? |
| [541] | 138 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | With results you mean evaluation results or sth different? |
| [542] | 139 | ch8-ai-assisted | alpha | must | done | so the thesis is considered as a kind of documentation or even specification? Ot |
| [543] | 139 | ch8-ai-assisted | alpha | must | done | ok - I understand the three first principles. But the last one was really applic |
| [544] | 139 | ch8-ai-assisted | alpha | must | done | to which category of memory systems it belongs? Please clarify. |
| [545] | 140 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | ok, so this information is complementary. I presume that this is needed for othe |
| [546] | 140 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | (between the agents) I assume that is the current case here ... |
| [547] | 141 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | ok but these are precise rates. As in the evaluation chapter, it was mentioned s |
| [548] | 141 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | what was wrong? Please provide your experience here ... |
| [549] | 141 | ch8-ai-assisted | alpha | must | done | Of course, I would expect that the thesis was written by you. And that the agent |
| [550] | 142 | ch8-ai-assisted | alpha | must | done | Ok, theoretically I see that there is a connection. It was also proven by the AI |
| [551] | 143 | ch8-ai-assisted | alpha | must | done | ok but this should have been evaluated more extensively to be sure about it ... |
| [552] | 144 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | returned? |
| [553] | 144 | ch8-ai-assisted | alpha | must | done | but was this the case? Did you also make manual changes in code in later commits |
| [554] | 145 | ch8-ai-assisted | alpha | must | addressed-by-rewrite | deadline pressures make artificial agents do mistakes? |
| [555] | 147 | ch9-discussion | beta-ref | should | addressed-by-rewrite | Please note that OpenAPI v2.X support was not examined in the evaluation. This i |
| [556] | 147 | ch9-discussion | generic | should | addressed-by-rewrite | threats to validity ... Are these also covered?  What is the structure of the ch |
| [557] | 147 | ch9-discussion | struct | must | addressed-by-rewrite | As I have already stated, please checks whether the term methodology that is too |
| [558] | 147 | ch9-discussion | typo | should | done | the part "..." |
| [559] | 147 | ch9-discussion | generic | should | addressed-by-rewrite | but functional tests don't they cover business logic errors? The API exposes an  |
| [560] | 148 | ch9-discussion | generic | should | addressed-by-rewrite | correct but it has to be mapped well, precisely and formally to these layers wit |
| [561] | 148 | ch9-discussion | generic | should | discussed | I believe that you did not comment on the part about quality gates.  In essence, |
| [562] | 148 | ch9-discussion | struct | must | addressed-by-rewrite | fully what? This research question was addressed fully? |
| [563] | 148 | ch9-discussion | typo | should | addressed-by-rewrite | OpenAPI |
| [564] | 148 | ch9-discussion | generic | defer | deferred | correct - but the business logic is already specified in the OpenAPI specificati |
| [565] | 149 | ch9-discussion | generic | should | addressed-by-rewrite | Could refer to the specific section with the large scale evaluation to be even m |
| [566] | 149 | ch9-discussion | generic | should | addressed-by-rewrite | could be removed + when this would take place? |
| [567] | 149 | ch9-discussion | generic | should | addressed-by-rewrite | Totally agree here. But I am not sure this was mentioned in the evaluation. So,  |
| [568] | 150 | ch9-discussion | generic | should | addressed-by-rewrite | but the finding seems to be negative while here you describe something positive  |
| [569] | 150 | ch9-discussion | generic | should | addressed-by-rewrite | logically speaking as part of a complex control-loop architecture with two level |
| [570] | 151 | ch9-discussion | generic | should | addressed-by-rewrite | did you forget security here? Otherwise, the math afterwards is imprecise as we  |
| [571] | 151 | ch9-discussion | typo | should | addressed-by-rewrite | its |
| [572] | 151 | ch9-discussion | typo | should | addressed-by-rewrite | the |
| [573] | 151 | ch9-discussion | generic | should | done | I agree with you here. Maybe you could add here that maybe the existing tooling  |
| [574] | 151 | ch9-discussion | generic | should | addressed-by-rewrite | at least the most critical ones (e.g., from staging to production). As you also  |
| [575] | 151 | ch9-discussion | beta-ref | should | addressed-by-rewrite | Please supply a respective reference here (to strengthen your argument/claim). |
| [576] | 152 | ch9-discussion | generic | should | discussed | Totally agree here. In fact, the developer could manually do the integration or  |
| [577] | 152 | ch9-discussion | generic | should | addressed-by-rewrite | However, this ... Ok but this is a valid scope restriction, especially in the co |
| [578] | 152 | ch9-discussion | generic | should | deferred | Correct! This was my own observation. In addition, did not evaluate whether all  |
| [579] | 152 | ch9-discussion | generic | should | addressed-by-rewrite | I agree here. The issue with the integration between an external Git service and |
| [580] | 152 | ch9-discussion | generic | should | discussed | or risk-based scoring model? It matches better the first option of severity-weig |
| [581] | 152 | ch9-discussion | generic | should | addressed-by-rewrite | sure but this is your own responsibility here. So, the question is whether it wa |
| [582] | 153 | ch9-discussion | generic | should | addressed-by-rewrite | correct - but still your system can also detect functional and non-functional is |
| [583] | 153 | ch9-discussion | generic | should | addressed-by-rewrite | ok but please decide. The benefit is for the platform team, the developer team o |
| [584] | 153 | ch9-discussion | generic | should | addressed-by-rewrite | ok - I do not disagree. But what is "meant" by model here? This has to be clarif |
| [585] | 155 | ch10-conclusion | struct | must | addressed-by-rewrite | again, please check the term as it is too strong for your contribution. |
| [586] | 155 | ch10-conclusion | typo | should | done | validity-mode-aware ... |
| [587] | 155 | ch10-conclusion | generic | should | addressed-by-rewrite | This is a ... which was also developed in the context of this thesis. |
| [588] | 155 | ch10-conclusion | typo | should | done | Based on this CRD, a ... |
| [589] | 156 | ch10-conclusion | generic | should | addressed-by-rewrite | This is a ... or This represents a ... |
| [590] | 156 | ch10-conclusion | generic | should | addressed-by-rewrite | ok but your answer does not cover the "comparable" part. Is it really comparable |
| [591] | 156 | ch10-conclusion | generic | should | addressed-by-rewrite | Again, "fully" is not sufficient - you need to explain what is "fully"! |
| [592] | 156 | ch10-conclusion | typo | should | addressed-by-rewrite | standard |
| [593] | 156 | ch10-conclusion | generic | should | addressed-by-rewrite | This was not so well advertised. In addition, I am not sure that it includes all |
| [594] | 156 | ch10-conclusion | generic | should | addressed-by-rewrite | Again, did not cover the "quality gate" part in the RQ ... |
| [595] | 157 | ch10-conclusion | generic | should | addressed-by-rewrite | The question is how effectively. So, your answer here should provide a degree .. |
| [596] | 157 | ch10-conclusion | generic | should | addressed-by-rewrite | OpenAPI specification? |
| [597] | 157 | ch10-conclusion | generic | should | addressed-by-rewrite | correct - did the evaluation proved that and how? For instance, does the reconci |
| [598] | 158 | ch10-conclusion | generic | should | addressed-by-rewrite | quality-gated pipelines can ... |
| [599] | 159 | ch10-conclusion | generic | should | addressed-by-rewrite | ok but is it sufficient? This is another major question here. I guess not. So, t |
| [600] | 159 | ch10-conclusion | generic | should | addressed-by-rewrite | Agreed. But even if the versioning strategy is described in the OpenAPI specific |
| [601] | 159 | ch10-conclusion | generic | should | addressed-by-rewrite | Plus the need to cover other kinds of non-functional testing, including security |
| [602] | 159 | ch10-conclusion | generic | should | addressed-by-rewrite | the other evaluation with different API versions is also limited ... |
| [603] | 159 | ch10-conclusion | generic | should | addressed-by-rewrite | I would also add the expressivity of the quality gate conditions - currently,  i |
| [604] | 159 | ch10-conclusion | generic | should | addressed-by-rewrite | I totally agree here. The question would be whether only this could be adjusted  |
| [605] | 159 | ch10-conclusion | generic | should | addressed-by-rewrite | Please highlight this as the main vision: having a closed control-loop potential |
| [606] | 160 | ch10-conclusion | generic | should | addressed-by-rewrite | Totally agree here. One additional thing: determining the right source of truth  |
| [607] | 160 | ch10-conclusion | generic | should | addressed-by-rewrite | ok, correct. But based on the promotion scenarios, adding these two principles i |
| [608] | 160 | ch10-conclusion | struct | must | addressed-by-rewrite | correct but the devil is in the details.  Would all principles be applicable to  |
| [609] | 160 | ch10-conclusion | generic | should | addressed-by-rewrite | Very nice scenario - quite realistic ...!!! |
| [610] | 160 | ch10-conclusion | generic | should | addressed-by-rewrite | So, the goal here is to create higher-level operators? I am asking this as in pr |

---

# Μέρος Γ — Σημεία συνειδητής διαφωνίας

Σε **10 σημεία** σκέφτηκα την πρότασή σας και επέλεξα να μην την ακολουθήσω. Δεν τα αγνόησα — κατέληξα σε διαφορετική απόφαση και είμαι έτοιμος να τα συζητήσω ζωντανά στην εξέταση.

## Γ.1 [65] / [556] — Διατήρηση Κεφαλαίου 9 ξεχωριστά από Κεφάλαιο 7

**Πρόταση Κρητικού:** Συγχώνευση Ch.9 (Discussion) με Ch.7 (Evaluation) — μόλις 6 σελίδες, θα μπορούσε να γίνει closing section του Ch.7.

**Η απόφασή μου:** Διατήρησα τον διαχωρισμό. Το Ch.7 αναφέρει αποδείξεις· το Ch.9 τις ερμηνεύει έναντι των RQs. Η συγχώνευση θα ανακάτευε «τι μετρήσαμε» με «τι σημαίνει αυτό». Αν προτιμάτε συγχώνευση για το bound copy, μπορώ να μετακινήσω το §9.5 Threats-to-Validity ως closing του Ch.7 και να κρατήσω το Ch.9 στις 4 σελίδες ερμηνείας.

## Γ.2 [557] — Διατήρηση τίτλου «Methodology: SDT» στο Ch.3

**Πρόταση Κρητικού:** «methodology» είναι υπερβολική λέξη, να γίνει «framework».

**Η απόφασή μου:** Round 1 ήδη αντικατέστησα τους περισσότερους «methodology» όρους με «framework». Διατήρησα όμως τον τίτλο του Ch.3 ως «Methodology: SDT» επειδή ο όρος εδώ αναφέρεται σε **research-process methodology** (πώς δομήθηκε η ερευνητική προσέγγιση), όχι σε **software-testing methodology** (formal test derivation procedure). Στο §3.1 το πρώτο εδάφιο εξηγεί τη διάκριση. Αν επιμένετε, μπορώ να το μετονομάσω σε «The SDT Framework».

## Γ.3 [447], [503] — Control-theory ως illustrative analogy μόνο

**Πρόταση Κρητικού:** Η control-theory ανάγνωση (stability, controllability, observability) είναι ισχυρή σύνδεση.

**Η απόφασή μου:** Συμφώνησα ως **illustrative analogy** και πρόσθεσα υποσημείωση που το διευκρινίζει. Δεν διεκδικώ formal Lyapunov stability ή formal observability — η αναλογία λειτουργεί απλώς για την περιγραφή του closed-loop στο GitOps domain. Είμαι έτοιμος να την αφαιρέσω αν τη βρίσκετε παραπλανητική.

## Γ.4 [115] — Διατήρηση «reconciliation» έναντι «adaptation»

**Πρόταση Κρητικού:** «Adaptation» καλύπτει καλύτερα τη σύγκλιση σε επιθυμητή κατάσταση.

**Η απόφασή μου:** Διατήρησα «reconciliation» γιατί είναι ο όρος-of-art στη Kubernetes/GitOps κοινότητα (Argo CD, Crossplane, CNCF OpenGitOps). Αυτή είναι η ορολογία που χρησιμοποιούν οι engineers που θα διαβάσουν τη διπλωματική. Το «adaptation» είναι ορθότερο σε control theory contexts αλλά λιγότερο αναγνωρίσιμο εδώ.

## Γ.5 [297] — P006 (Functional Testing) ως single-interaction

**Πρόταση Κρητικού:** Πραγματικά APIs χρειάζονται multi-step sequences και stateful testing.

**Η απόφασή μου:** Αναγνωρίζω ότι το P006 του DriveBy εκτελεί single-shot schema-derived requests. Stateful testing είναι **named future work** στο §10.5 (με αναφορά σε RESTler, EvoMaster ως υφιστάμενες υλοποιήσεις). Η σχεδιαστική επιλογή για single-shot είναι έντονη: P006 ήταν η πρώτη runtime principle, και η μεθοδολογική θεμελίωση είχε προτεραιότητα έναντι της εκτεταμένης λειτουργικότητας.

## Γ.6 [580] — Severity-weighted scoring (όχι risk-based)

**Πρόταση Κρητικού:** Risk-based scoring θα ήταν πιο αξιόλογο.

**Η απόφασή μου:** Το severity-weighted scoring είναι το **απλούστερο** μοντέλο που χαρτογραφεί σε quality gate (critical → block, warning → record-but-pass). Risk-based scoring (severity × likelihood × blast radius × fix cost) απαιτεί empirical calibration data που η διπλωματική δεν συλλέγει. Καταγράφηκε ως μελλοντική εργασία στο §10.4 με συγκεκριμένο hook (ο JSON output ήδη φέρει τα απαραίτητα structured fields).

## Γ.7 [246] — Διατήρηση τίτλου «DriveBy CLI Architecture» (Ch.4)

**Πρόταση Κρητικού:** Καλύτερα «The DriveBy System» αντί για «CLI Architecture».

**Η απόφασή μου:** Διατήρησα «DriveBy CLI Architecture» επειδή το Ch.5 ήδη καλύπτει το ευρύτερο σύστημα (Kubernetes operator, XSDLC, GitOps quality gate). Το Ch.4 είναι ειδικά για το binary του CLI: το dependency flow, το APISpec abstraction, το principle-checker pattern, το engine. Αν προτιμάτε «The DriveBy CLI» αντί «DriveBy CLI Architecture», είναι ευχερές edit.

## Γ.8 [140] — Διατήρηση «Static Analysis» (όχι «Syntactic Validation»)

**Πρόταση Κρητικού:** «Syntactic Validation» είναι ακριβέστερος όρος.

**Η απόφασή μου:** Διατήρησα «Static Analysis» γιατί είναι ο field-standard term για αυτά που κάνουν Spectral/Vacuum/Redocly (rule-based linting σε structured documents). Το «Syntactic Validation» θα στενεύει την κατηγορία — οι Spectral rules καλύπτουν documentation completeness και security-scheme presence, όχι μόνο syntax. Στο §2.6 πρόσθεσα terminological note που αναγνωρίζει την παρατήρησή σας.

## Γ.9 P006/P007 ως planned PrincipleCheckers (όχι full)

**Πρόταση Κρητικού (μη ρητή, αλλά αναμενόμενη):** Γιατί 7/9 full PrincipleCheckers και 2 test runners;

**Η απόφασή μου:** Οι runtime principles (P006, P007) είναι I/O-bound (HTTP client, target host, timeout). Οι static principles είναι pure functions. Η current architecture εκθέτει τις runtime principles ως ξεχωριστούς test runners που παράγουν το ίδιο per-principle JSON output όπως η PrincipleChecker registry — οπότε downstream consumers δεν βλέπουν διαφορά. Το wrapper integration είναι **queued mechanical refactor**, όχι νέα σχεδίαση. Στο §3.3 ο πίνακας metadata δηλώνει ρητά Implementation Status: P006/P007 = «Defined theoretically; PrincipleChecker wrapper pending».

## Γ.10 Single-cluster, single-team PoC

**Πρόταση Κρητικού (μη ρητή):** Πόσοι developers χρησιμοποίησαν πραγματικά το σύστημα;

**Η απόφασή μου:** Το PoC είναι **σκόπιμα μικρό**: ένας cluster (`private.novelcore.org`), ένας platform-team operator (ο συγγραφέας), 5 evaluation APIs, εκατοντάδες gate executions. Δεν διεκδικώ multi-team deployment — το §7.5 περιγράφει ακριβώς τι υπήρχε πριν το PoC και τι προσέθεσε. Η εναλλακτική θα ήταν synthetic minikube demo που δεν θα έδειχνε την **platform-inheritance property** (§5.2, §7.5). Multi-cluster, multi-team scaling είναι named future work στο §10.5.

---

# Συνοπτικά: Headline metrics μεταξύ rounds

| Metric | Round 0 (παράδοση) | Round 10 (τρέχουσα) |
|---|---|---|
| Σελίδες | 165 | **216** |
| Σχόλια supervisor | 610 | — |
| Σχόλια `open` ή `needs-walkthrough` | 610 | **0** |
| Σχόλια `done` (ρητά διορθωμένα) | 0 | **75** |
| Σχόλια `addressed-by-rewrite` | 0 | **505** |
| Σχόλια `discussed` | 0 | **23** |
| Σχόλια `deferred` | 0 | **7** |
| LaTeX warnings | πολλά overfull | **0** |
| Undefined cross-refs | αρκετά | **0** |
| Νέα διαγράμματα | 0 | **3** |
| Νέα evaluation arms | 0 | **4** |
| Νέα reproducible artefacts | 0 | **3 INDEX.md folders** |
| Front matter (EN + GR abstract + LoF + LoT + acronyms) | μερικό | **πλήρες, bilingual** |

---

# Παραρτήματα

- `thesis/main.pdf` — η αναθεωρημένη διπλωματική (216 σελ., 0 overfull, 0 errors).
- `thesis/review/kritikos-comments.md` — το extract και των 610 σχολίων από το annotated PDF.
- `thesis/review/comment-index.md` — ο πλήρης πίνακας 610 γραμμών με per-row trailer `[r4: ...]` που ονομάζει την ενότητα που απαντά στο σχόλιο.
- `thesis/review/changes-log.md` — log αλλαγών με commit/section references.
- `thesis/review/defense-qa.md` — προετοιμασμένες απαντήσεις για τα 10 σημεία διαφωνίας.
- `results/sdt-feedback/INDEX.md` — artefacts του SDT-feedback πειράματος ([550]).
- `results/per-check-injection/INDEX.md` — artefacts του per-check defect injection ([551], [463]).
- `results/apisguru-rerun/INDEX.md` — artefacts της APIs.guru rerun ([485]–[490]).

Σας ευχαριστώ ξανά για τον χρόνο και την προσοχή που αφιερώσατε στην ανατροφοδότηση. Η διπλωματική είναι ουσιαστικά καλύτερη χάρη στην ποιότητα και το βάθος των σχολίων σας. Είμαι στη διάθεσή σας για περαιτέρω συζήτηση πριν την εξέταση τον Ιούνιο.

Με εκτίμηση,

Πέτρος Ευάγγελος Τριανταφύλλης

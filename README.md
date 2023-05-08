# Govulnapi

Govulnapi is a deliberately vulnerable API written in Go showcasing common security flaws that are made while developing an application backend. Govulnapi aims be be concise in respect to vulnerabilities it implements as to lower the entry barrier for junior security researchers (e.g. "Use of Weak Hash" weakness can be found in the codebase by searching for a comment containing its corresponding id CWE-328).

## Setup

### With docker

1. `git clone --depth 1 https://github.com/govulnapi/govulnapi.git`
2. `cd govulnapi`
3. `make build`
4. `make run`

## Servers

- Web client: <http://localhost:8080/>
- API documentation: <http://localhost:8081/>
- Virtual Coingecko: <http://localhost:8082/>

## Implemented vulnerabilities

### [OWASP Top 10 2023 - draft](https://github.com/OWASP/API-Security/tree/master/2023/en/src) (TBD)

### [OWASP Top 10 2021](https://owasp.org/www-project-top-ten)

- [ ] [A01 - Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control)

  - [x] [CWE-276: Incorrect Default Permissions](https://cwe.mitre.org/data/definitions/276.html)
  - [x] [CWE-639: Authorization Bypass Through User-Controlled Key](https://cwe.mitre.org/data/definitions/639.html)
  - [x] [CWE-942: Permissive Cross-domain Policy with Untrusted Domains](https://cwe.mitre.org/data/definitions/942.html)

- [ ] [A02 - Cryptographic Failures](https://owasp.org/Top10/A02_2021-Cryptographic_Failures)

  - [x] [CWE-327: Use of a Broken or Risky Cryptographic Algorithm](https://cwe.mitre.org/data/definitions/327.html)
  - [x] [CWE-319: Cleartext Transmission of Sensitive Information](https://cwe.mitre.org/data/definitions/319.html)
  - [x] [CWE-328: Use of Weak Hash](https://cwe.mitre.org/data/definitions/328.html)
  - [x] [CWE-340: Generation of Predictable Numbers or Identifiers](https://cwe.mitre.org/data/definitions/340.html)
  - [x] [CWE-523: Unprotected Transport of Credentials](https://cwe.mitre.org/data/definitions/523.html)
  - [x] [CWE-759: Use of a One-Way Hash without a Salt](https://cwe.mitre.org/data/definitions/759.html)
  - [x] [CWE-916: Use of Password Hash With Insufficient Computational Effort](https://cwe.mitre.org/data/definitions/916.html)

- [ ] [A03 - Injection](https://owasp.org/Top10/A03_2021-Injection)

  - [x] [CWE-20: Improper Input Validation](https://cwe.mitre.org/data/definitions/20.html)
  - [x] [CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')](https://cwe.mitre.org/data/definitions/79.html)
  - [x] [CWE-89: Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')](https://cwe.mitre.org/data/definitions/89.html)

- [ ] [A04 - Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design)

  - [x] [CWE-256: Plaintext Storage of a Password](https://cwe.mitre.org/data/definitions/256.html)
  - [x] [CWE-472: External Control of Assumed-Immutable Web Parameter](https://cwe.mitre.org/data/definitions/472.html)
  - [x] [CWE-598: Use of GET Request Method With Sensitive Query Strings](https://cwe.mitre.org/data/definitions/598.html)

- [ ] [A05 - Security Misconfiguration](https://owasp.org/Top10/A05_2021-Security_Misconfiguration)

  - [x] [CWE-547: Use of Hard-coded, Security-relevant Constants](https://cwe.mitre.org/data/definitions/547.html)
  - [x] [CWE-614: Sensitive Cookie in HTTPS Session Without 'Secure' Attribute](https://cwe.mitre.org/data/definitions/614.html)
  - [x] [CWE-1004: Sensitive Cookie Without 'HttpOnly' Flag](https://cwe.mitre.org/data/definitions/1004.html)

- [x] [A06 - Vulnerable and Outdated Components](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components)

  - [x] [CWE-1104: Use of Unmaintained Third Party Components](https://cwe.mitre.org/data/definitions/1104.html)

- [ ] [A07 - Identification and Authentication Failures](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures)

  - [x] [CWE-262: Not Using Password Aging](https://cwe.mitre.org/data/definitions/262.html)
  - [x] [CWE-521: Weak Password Requirements](https://cwe.mitre.org/data/definitions/521.html)
  - [x] [CWE-549: Missing Password Field Masking](https://cwe.mitre.org/data/definitions/549.html)
  - [x] [CWE-620: Unverified Password Change](https://cwe.mitre.org/data/definitions/620.html)

- [ ] [A08 - Software and Data Integrity Failures](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures)

  - [x] [CWE-613: Insufficient Session Expiration](https://cwe.mitre.org/data/definitions/613.html)
  - [x] [CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes](https://cwe.mitre.org/data/definitions/502.html)

- [ ] [A09 - Security Logging and Monitoring Failures](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures)

  - [x] [CWE-223: Omission of Security-relevant Information](https://cwe.mitre.org/data/definitions/223.html)
  - [x] [CWE-532: Insertion of Sensitive Information into Log File](https://cwe.mitre.org/data/definitions/532.html)
  - [x] [CWE-778: Insufficient Logging](https://cwe.mitre.org/data/definitions/778.html)

- [ ] [A10 - Server-Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29)
  - [ ] [CWE-441: Unintended Proxy or Intermediary ('Confused Deputy')](https://cwe.mitre.org/data/definitions/441.html)

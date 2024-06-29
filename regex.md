| Key | Value | Description |
| --- | --- | --- |
| `^` | Start of line | Matches the beginning of a line. |
| `$` | End of line | Matches the end of a line. |
| `.` | Any character | Matches any single character. |
| `\w` | Word character | Matches any word character (alphanumeric plus underscore). |
| `\W` | Non-word character | Matches any non-word character. |
| `[^)]*` | Group with any character | Matches any character not in a character class. |
| `*` | Zero or more occurrences | Matches zero or more occurrences of the preceding pattern. |
| `+` | One or more occurrences | Matches one or more occurrences of the preceding pattern. |
| `?` | Zero or one occurrence | Matches zero or one occurrence of the preceding pattern. |
| `{}` | Exactly n occurrences | Matches exactly n occurrences of the preceding pattern. |
| `|` | Alternative | Matches either the preceding or following pattern. |

**Character Classes**

| Key | Value | Description |
| --- | --- | --- |
| `[]` | Character class | Matches any character in the specified class. |
| `\` | Backslash | Escape special characters. |
| `^` | Caret | Matches the beginning of a string. |
| `$` | Dollar sign | Matches the end of a string. |
| `\w` | Word character | Matches any word character (alphanumeric plus underscore). |
| `\W` | Non-word character | Matches any non-word character. |
| `]` | Close bracket | Closes an open character class. |
| `^` | Start of line | Matches the beginning of a line. |
| `$` | End of line | Matches the end of a line. |

**Quantifiers**

| Key | Value | Description |
| --- | --- | --- |
| `*` | Zero or more occurrences | Matches zero or more occurrences of the preceding pattern. |
| `+` | One or more occurrences | Matches one or more occurrences of the preceding pattern. |
| `?` | Zero or one occurrence | Matches zero or one occurrence of the preceding pattern. |
| `{n, m}` | Exactly n to m occurrences | Matches exactly n to m occurrences of the preceding pattern. |

**Advanced Patterns**

| Key | Value | Description |
| --- | --- | --- |
| `\A` | Beginning of line | Matches the beginning of a line. |
| `\Z` | End of line | Matches the end of a line. |
| `\b` | Word boundary | Matches a word boundary (the start or end of a word). |
| `\B` | Non-word boundary | Matches any character that is not at a word boundary. |
| `\w{3,}` | Exactly 3 to 5 characters long | Matches exactly 3 to 5 alphanumeric characters. |
| `\W{3,}` | Exactly 3 to 5 non-word characters | Matches exactly 3 to 5 non-alphanumeric characters. |
| `(?:pattern)` | Non-capturing group | Creates a non-capturing group around the preceding pattern. |
| `(?(group)pattern)` | Capturing group | Creates a capturing group around the preceding pattern. |
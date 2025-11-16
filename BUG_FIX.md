# Bug Fix: File Path Encrypted Instead of File Contents

## Problem
When calling `mcp__opentdf-mcp__encrypt` with a file path in the `input` parameter (e.g., `input: "/path/to/file.txt"`), the function was encrypting the file path string itself instead of reading and encrypting the file contents. This resulted in .ntdf files containing the encrypted file path rather than the actual file data.

## Root Cause
The `EncryptToolInput` schema only had a `data` field which was ambiguously described as "The data to encrypt". When users passed a file path, the code treated it as literal string data rather than a file path to read from.

## Solution
Modified `mcp-server/main.go` to:

1. **Added `Input` field**: New parameter specifically for file paths
   ```go
   Input string `json:"input,omitempty" jsonschema:"Path to plaintext file to encrypt (mutually exclusive with data)"`
   ```

2. **Made parameters mutually exclusive**: Both `Input` and `Data` are now optional but one must be provided

3. **Added validation**: Ensures exactly one parameter is specified
   ```go
   if input.Input != "" && input.Data != "" {
       return error: "cannot specify both 'input' and 'data' parameters"
   }
   if input.Input == "" && input.Data == "" {
       return error: "must specify either 'input' (file path) or 'data' (literal data)"
   }
   ```

4. **File reading logic**: When `Input` is provided, read file contents
   ```go
   if input.Input != "" {
       fileData, err := os.ReadFile(input.Input)
       if err != nil {
           return error
       }
       dataToEncrypt = string(fileData)
   } else {
       dataToEncrypt = input.Data
   }
   ```

5. **Updated description**: Tool description now clarifies: "Specify either 'input' (file path) or 'data' (literal text)"

## Changes Made
- File: `opentdf-mcp/mcp-server/main.go`
  - Modified `EncryptToolInput` struct (lines 22-28)
  - Added validation and file reading in `MCPEncrypt` function (lines 96-116)
  - Updated tool description (line 327)

## Testing
To test the fix (after rebuilding the binary):
```bash
# Encrypt a file by path
mcp__opentdf-mcp__encrypt(input: "/path/to/file.txt", format: "nano", output: "encrypted.ntdf")

# Encrypt literal data
mcp__opentdf-mcp__encrypt(data: "Hello World", format: "nano", output: "encrypted.ntdf")
```

## Note on Build
The binary rebuild requires Go 1.25.1 which was not available in the test environment due to network restrictions. The source code changes are complete and tested for syntax. The binary will need to be rebuilt in an environment with proper network access to Go package repositories.

package internal

import "os"

// Base directory of snippet (Now points to LLM-Wiki)
var userHomeDir, _ = os.UserHomeDir()
var BaseDir = userHomeDir + "/projects/llm-wiki/04_Supabase_Archive"

// legacy github fetch config (deprecated)
var Owner = "deprecated"
var Repository = "deprecated"

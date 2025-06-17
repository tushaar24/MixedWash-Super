package config

import (

	"log"

	supabase "github.com/supabase-community/supabase-go"
)

func GetSupabaseClient() *supabase.Client {
	const supabaseUrl = "https://ainlzqfaijqluyuqdojj.supabase.co"
	const supabaseKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFpbmx6cWZhaWpxbHV5dXFkb2pqIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTc0NzM5MzIwMywiZXhwIjoyMDYyOTY5MjAzfQ.liqyqsttU3o2IgwhdwCHnUAdb7KNgNXWJlveslH3loI"

	client, err := supabase.NewClient(supabaseUrl, supabaseKey, &supabase.ClientOptions{},)
	if err != nil {
		log.Fatalf("Failed to initialise supabase client: %v", err)
	}
	return client
}

```proto
syntax = "proto3";

package go.lsp.dev.textDocument;

option java_multiple_files = true;

option go_package = "go.lsp.dev.textDocument;textDocument";

option cc_enable_arenas = true;

option csharp_namespace = "Go.Lsp.Dev.TextDocument";

option java_package = "dev.lsp.go";

option java_outer_classname = "TextDocument";

message Range {
  Position end = 1;

  Position start = 2;
}

message TextDocument {
  string language_id = 1;

  int32 line_count = 2;

  DocumentUri uri = 3;

  int32 version = 4;
}

message TextDocumentContentChangeEvent {
  oneof text_document_content_change_event {
    TextDocumentContentChangeEvent_0 text_document_content_change_event_0 = 1;

    TextDocumentContentChangeEvent_1 text_document_content_change_event_1 = 2;
  }

  message TextDocumentContentChangeEvent_0 {
    Range range = 1;

    int32 range_length = 2;

    string text = 3;
  }

  message TextDocumentContentChangeEvent_1 {
    string text = 1;
  }
}

message TextEdit {
  string new_text = 1;

  Range range = 2;
}

message DocumentUri {
  string document_uri = 1;
}

message FullTextDocument {
  string _language_id = 1;

  string uri = 2;

  int32 version = 3;

  int32 _version = 4;

  string language_id = 5;

  int32 line_count = 6;

  string _content = 7;

  repeated int32 _line_offsets = 8;

  DocumentUri _uri = 9;
}

message Position {
  int32 character = 1;

  int32 line = 2;
}
```

```proto
syntax = "proto3";

package go.lsp.dev.textDocument;

option java_package = "dev.lsp.go";

option java_outer_classname = "TextDocument";

option java_multiple_files = true;

option go_package = "go.lsp.dev.textDocument;textDocument";

option cc_enable_arenas = true;

option csharp_namespace = "Go.Lsp.Dev.TextDocument";

message AnnotatedTextEdit {
  string new_text = 1;

  Range range = 2;

  ChangeAnnotationIdentifier annotation_id = 3;
}

message DocumentUri {
  string document_uri = 1;
}

message Decimal {
  int32 decimal = 1;
}

message Integer {
  int32 integer = 1;
}

message Uinteger {
  int32 uinteger = 1;
}

message ChangeAnnotationIdentifier {
  string change_annotation_identifier = 1;
}

message Position {
  Uinteger character = 1;

  Uinteger line = 2;
}

message Range {
  Position end = 1;

  Position start = 2;
}

enum SymbolKind {
  SymbolKind_1 = 1;

  SymbolKind_2 = 2;

  SymbolKind_3 = 3;

  SymbolKind_4 = 4;

  SymbolKind_5 = 5;

  SymbolKind_6 = 6;

  SymbolKind_7 = 7;

  SymbolKind_8 = 8;

  SymbolKind_9 = 9;

  SymbolKind_10 = 10;

  SymbolKind_11 = 11;

  SymbolKind_12 = 12;

  SymbolKind_13 = 13;

  SymbolKind_14 = 14;

  SymbolKind_15 = 15;

  SymbolKind_16 = 16;

  SymbolKind_17 = 17;

  SymbolKind_18 = 18;

  SymbolKind_19 = 19;

  SymbolKind_20 = 20;

  SymbolKind_21 = 21;

  SymbolKind_22 = 22;

  SymbolKind_23 = 23;

  SymbolKind_24 = 24;

  SymbolKind_25 = 25;

  SymbolKind_26 = 26;
}
```

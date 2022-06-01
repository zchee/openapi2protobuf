```proto
syntax = "proto3";

package go.lsp.dev.textDocument;

option java_package = "dev.lsp.go";

option java_outer_classname = "TextDocument";

option java_multiple_files = true;

option go_package = "go.lsp.dev.textDocument;textDocument";

option cc_enable_arenas = true;

option csharp_namespace = "Go.Lsp.Dev.TextDocument";

message SymbolKind {
  enum SymbolKind {
    SymbolKind_1 = 1;

    SymbolKind_10 = 2;

    SymbolKind_11 = 3;

    SymbolKind_12 = 4;

    SymbolKind_13 = 5;

    SymbolKind_14 = 6;

    SymbolKind_15 = 7;

    SymbolKind_16 = 8;

    SymbolKind_17 = 9;

    SymbolKind_18 = 10;

    SymbolKind_19 = 11;

    SymbolKind_2 = 12;

    SymbolKind_20 = 13;

    SymbolKind_21 = 14;

    SymbolKind_22 = 15;

    SymbolKind_23 = 16;

    SymbolKind_24 = 17;

    SymbolKind_25 = 18;

    SymbolKind_26 = 19;

    SymbolKind_3 = 20;

    SymbolKind_4 = 21;

    SymbolKind_5 = 22;

    SymbolKind_6 = 23;

    SymbolKind_7 = 24;

    SymbolKind_8 = 25;

    SymbolKind_9 = 26;
  }
}

message Integer {
  int32 integer = 1;
}

message ChangeAnnotation {
  string label = 1;

  bool needs_confirmation = 2;

  string description = 3;
}

message ChangeAnnotations {
  Annotations _annotations = 1;

  int32 _counter = 2;

  int32 _size = 3;

  int32 size = 4;

  message Annotations {
  }
}

message DocumentUri {
  string document_uri = 1;
}

message LSPArray {
  repeated  lsp_any = 1;
}

message Range {
  Position end = 1;

  Position start = 2;
}

message Position {
  Uinteger character = 1;

  Uinteger line = 2;
}

message AnnotatedTextEdit {
  ChangeAnnotationIdentifier annotation_id = 1;

  string new_text = 2;

  Range range = 3;
}

message  {
  enum  {
    _1 = 1;
  }
}

message LSPAny1 {
  LSPObject lsp_object = 1;

  LSPArray lsp_array = 2;

  message LSPObject {
  }

  message LSPArray {
    repeated  lsp_any = 1;
  }
}

message CallHierarchyIncomingCall {
  CallHierarchyItem from = 1;

  repeated FromRanges from_ranges = 2;

  message FromRanges {
    repeated  range = 1;
  }
}

message CallHierarchyOutgoingCall {
  repeated FromRanges from_ranges = 1;

  CallHierarchyItem to = 2;

  message FromRanges {
    repeated  range = 1;
  }
}

message ChangeAnnotationIdentifier1 {
  string change_annotation_identifier_1 = 1;
}

message LSPAny {
  LSPObject lsp_object = 1;

    = 2;

  message LSPObject {
  }

  message  {
    repeated  lsp_any = 1;
  }
}

message CallHierarchyItem {
  LSPAny1 data = 1;

  string detail = 2;

  SymbolKind kind = 3;

  string name = 4;

  Range range = 5;

  Range selection_range = 6;

  repeated Tags tags = 7;

  DocumentUri uri = 8;

  message Tags {
     tags = 1;

    message  {
      enum  {
        _1 = 1;
      }
    }
  }
}

message ChangeAnnotationIdentifier {
  string change_annotation_identifier = 1;
}

message Uinteger {
  int32 uinteger = 1;
}
```

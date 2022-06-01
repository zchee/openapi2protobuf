```proto
name: "text_document"
package: "go.lsp.dev.textDocument"
message_type: {
  name: "BaseSymbolInformation"
  field: {
    name: "container_name"
    number: 1
    type: TYPE_STRING
    type_name: "TYPE_STRING"
  }
  field: {
    name: "name"
    number: 2
    type: TYPE_STRING
    type_name: "TYPE_STRING"
  }
}
message_type: {
  name: "Declaration"
  field: {
    name: "declaration_1"
    number: 1
    type: TYPE_MESSAGE
    type_name: "Declaration_1"
    oneof_index: 0
  }
  field: {
    name: "declaration_2"
    number: 2
    type: TYPE_MESSAGE
    type_name: "Declaration_2"
    oneof_index: 0
  }
  nested_type: {
    name: "Declaration_1"
    field: {
      name: "uri"
      number: 1
      type: TYPE_MESSAGE
      type_name: "DocumentUri"
    }
    field: {
      name: "range"
      number: 2
      type: TYPE_MESSAGE
      type_name: "Range"
    }
  }
  nested_type: {
    name: "Declaration_2"
    field: {
      name: "location"
      number: 1
      label: LABEL_REPEATED
      type: TYPE_MESSAGE
      type_name: "Location"
    }
    nested_type: {
      name: "Location"
      field: {
        name: "range"
        number: 1
        type: TYPE_MESSAGE
        type_name: "Range"
      }
      field: {
        name: "uri"
        number: 2
        type: TYPE_MESSAGE
        type_name: "DocumentUri"
      }
    }
  }
  oneof_decl: {
    name: "declaration"
  }
}
message_type: {
  name: "ChangeAnnotations"
  field: {
    name: "_counter"
    number: 1
    type: TYPE_INT32
    type_name: "TYPE_INT32"
  }
  field: {
    name: "_size"
    number: 2
    type: TYPE_INT32
    type_name: "TYPE_INT32"
  }
  field: {
    name: "size"
    number: 3
    type: TYPE_INT32
    type_name: "TYPE_INT32"
  }
}
message_type: {
  name: "DocumentUri"
  field: {
    name: "document_uri"
    number: 1
    type: TYPE_STRING
  }
}
message_type: {
  name: "Position"
  field: {
    name: "character"
    number: 1
    type: TYPE_MESSAGE
    type_name: "Uinteger"
  }
  field: {
    name: "line"
    number: 2
    type: TYPE_MESSAGE
    type_name: "Uinteger"
  }
}
message_type: {
  name: "ChangeAnnotation"
  field: {
    name: "description"
    number: 1
    type: TYPE_STRING
    type_name: "TYPE_STRING"
  }
  field: {
    name: "label"
    number: 2
    type: TYPE_STRING
    type_name: "TYPE_STRING"
  }
  field: {
    name: "needs_confirmation"
    number: 3
    type: TYPE_BOOL
    type_name: "TYPE_BOOL"
  }
}
message_type: {
  name: "ChangeAnnotationIdentifier"
  field: {
    name: "change_annotation_identifier"
    number: 1
    type: TYPE_STRING
  }
}
message_type: {
  name: "Uinteger"
  field: {
    name: "uinteger"
    number: 1
    type: TYPE_INT32
  }
}
message_type: {
  name: "Location"
  field: {
    name: "range"
    number: 1
    type: TYPE_MESSAGE
    type_name: "Range"
  }
  field: {
    name: "uri"
    number: 2
    type: TYPE_MESSAGE
    type_name: "DocumentUri"
  }
}
message_type: {
  name: "Range"
  field: {
    name: "start"
    number: 1
    type: TYPE_MESSAGE
    type_name: "Position"
  }
  field: {
    name: "end"
    number: 2
    type: TYPE_MESSAGE
    type_name: "Position"
  }
}
message_type: {
  name: "Decimal"
  field: {
    name: "decimal"
    number: 1
    type: TYPE_INT32
  }
}
message_type: {
  name: "Integer"
  field: {
    name: "integer"
    number: 1
    type: TYPE_INT32
  }
}
message_type: {
  name: "AnnotatedTextEdit"
  field: {
    name: "annotation_id"
    number: 1
    type: TYPE_MESSAGE
    type_name: "ChangeAnnotationIdentifier"
  }
  field: {
    name: "new_text"
    number: 2
    type: TYPE_STRING
    type_name: "TYPE_STRING"
  }
  field: {
    name: "range"
    number: 3
    type: TYPE_MESSAGE
    type_name: "Range"
  }
}
enum_type: {
  name: "SemanticTokenTypes"
  value: {
    name: "SemanticTokenTypes_Class"
    number: 1
  }
  value: {
    name: "SemanticTokenTypes_Comment"
    number: 2
  }
  value: {
    name: "SemanticTokenTypes_Decorator"
    number: 3
  }
  value: {
    name: "SemanticTokenTypes_Enum"
    number: 4
  }
  value: {
    name: "SemanticTokenTypes_EnumMember"
    number: 5
  }
  value: {
    name: "SemanticTokenTypes_Event"
    number: 6
  }
  value: {
    name: "SemanticTokenTypes_Function"
    number: 7
  }
  value: {
    name: "SemanticTokenTypes_Interface"
    number: 8
  }
  value: {
    name: "SemanticTokenTypes_Keyword"
    number: 9
  }
  value: {
    name: "SemanticTokenTypes_Macro"
    number: 10
  }
  value: {
    name: "SemanticTokenTypes_Method"
    number: 11
  }
  value: {
    name: "SemanticTokenTypes_Modifier"
    number: 12
  }
  value: {
    name: "SemanticTokenTypes_Namespace"
    number: 13
  }
  value: {
    name: "SemanticTokenTypes_Number"
    number: 14
  }
  value: {
    name: "SemanticTokenTypes_Operator"
    number: 15
  }
  value: {
    name: "SemanticTokenTypes_Parameter"
    number: 16
  }
  value: {
    name: "SemanticTokenTypes_Property"
    number: 17
  }
  value: {
    name: "SemanticTokenTypes_Regexp"
    number: 18
  }
  value: {
    name: "SemanticTokenTypes_String"
    number: 19
  }
  value: {
    name: "SemanticTokenTypes_Struct"
    number: 20
  }
  value: {
    name: "SemanticTokenTypes_Type"
    number: 21
  }
  value: {
    name: "SemanticTokenTypes_TypeParameter"
    number: 22
  }
  value: {
    name: "SemanticTokenTypes_Variable"
    number: 23
  }
}
enum_type: {
  name: "SymbolKind"
  value: {
    name: "SymbolKind_1"
    number: 1
  }
  value: {
    name: "SymbolKind_2"
    number: 2
  }
  value: {
    name: "SymbolKind_3"
    number: 3
  }
  value: {
    name: "SymbolKind_4"
    number: 4
  }
  value: {
    name: "SymbolKind_5"
    number: 5
  }
  value: {
    name: "SymbolKind_6"
    number: 6
  }
  value: {
    name: "SymbolKind_7"
    number: 7
  }
  value: {
    name: "SymbolKind_8"
    number: 8
  }
  value: {
    name: "SymbolKind_9"
    number: 9
  }
  value: {
    name: "SymbolKind_10"
    number: 10
  }
  value: {
    name: "SymbolKind_11"
    number: 11
  }
  value: {
    name: "SymbolKind_12"
    number: 12
  }
  value: {
    name: "SymbolKind_13"
    number: 13
  }
  value: {
    name: "SymbolKind_14"
    number: 14
  }
  value: {
    name: "SymbolKind_15"
    number: 15
  }
  value: {
    name: "SymbolKind_16"
    number: 16
  }
  value: {
    name: "SymbolKind_17"
    number: 17
  }
  value: {
    name: "SymbolKind_18"
    number: 18
  }
  value: {
    name: "SymbolKind_19"
    number: 19
  }
  value: {
    name: "SymbolKind_20"
    number: 20
  }
  value: {
    name: "SymbolKind_21"
    number: 21
  }
  value: {
    name: "SymbolKind_22"
    number: 22
  }
  value: {
    name: "SymbolKind_23"
    number: 23
  }
  value: {
    name: "SymbolKind_24"
    number: 24
  }
  value: {
    name: "SymbolKind_25"
    number: 25
  }
  value: {
    name: "SymbolKind_26"
    number: 26
  }
}
enum_type: {
  name: "Tags"
  value: {
    name: "Tags_1"
    number: 1
  }
}
enum_type: {
  name: "InlayHintKind"
  value: {
    name: "InlayHintKind_1"
    number: 1
  }
  value: {
    name: "InlayHintKind_2"
    number: 2
  }
}
enum_type: {
  name: "SemanticTokenModifiers"
  value: {
    name: "SemanticTokenModifiers_Abstract"
    number: 1
  }
  value: {
    name: "SemanticTokenModifiers_Async"
    number: 2
  }
  value: {
    name: "SemanticTokenModifiers_Declaration"
    number: 3
  }
  value: {
    name: "SemanticTokenModifiers_DefaultLibrary"
    number: 4
  }
  value: {
    name: "SemanticTokenModifiers_Definition"
    number: 5
  }
  value: {
    name: "SemanticTokenModifiers_Deprecated"
    number: 6
  }
  value: {
    name: "SemanticTokenModifiers_Documentation"
    number: 7
  }
  value: {
    name: "SemanticTokenModifiers_Modification"
    number: 8
  }
  value: {
    name: "SemanticTokenModifiers_Readonly"
    number: 9
  }
  value: {
    name: "SemanticTokenModifiers_Static"
    number: 10
  }
}
enum_type: {
  name: "DocumentHighlightKind"
  value: {
    name: "DocumentHighlightKind_1"
    number: 1
  }
  value: {
    name: "DocumentHighlightKind_2"
    number: 2
  }
  value: {
    name: "DocumentHighlightKind_3"
    number: 3
  }
}
enum_type: {
  name: "CompletionItemKind"
  value: {
    name: "CompletionItemKind_1"
    number: 1
  }
  value: {
    name: "CompletionItemKind_2"
    number: 2
  }
  value: {
    name: "CompletionItemKind_3"
    number: 3
  }
  value: {
    name: "CompletionItemKind_4"
    number: 4
  }
  value: {
    name: "CompletionItemKind_5"
    number: 5
  }
  value: {
    name: "CompletionItemKind_6"
    number: 6
  }
  value: {
    name: "CompletionItemKind_7"
    number: 7
  }
  value: {
    name: "CompletionItemKind_8"
    number: 8
  }
  value: {
    name: "CompletionItemKind_9"
    number: 9
  }
  value: {
    name: "CompletionItemKind_10"
    number: 10
  }
  value: {
    name: "CompletionItemKind_11"
    number: 11
  }
  value: {
    name: "CompletionItemKind_12"
    number: 12
  }
  value: {
    name: "CompletionItemKind_13"
    number: 13
  }
  value: {
    name: "CompletionItemKind_14"
    number: 14
  }
  value: {
    name: "CompletionItemKind_15"
    number: 15
  }
  value: {
    name: "CompletionItemKind_16"
    number: 16
  }
  value: {
    name: "CompletionItemKind_17"
    number: 17
  }
  value: {
    name: "CompletionItemKind_18"
    number: 18
  }
  value: {
    name: "CompletionItemKind_19"
    number: 19
  }
  value: {
    name: "CompletionItemKind_20"
    number: 20
  }
  value: {
    name: "CompletionItemKind_21"
    number: 21
  }
  value: {
    name: "CompletionItemKind_22"
    number: 22
  }
  value: {
    name: "CompletionItemKind_23"
    number: 23
  }
  value: {
    name: "CompletionItemKind_24"
    number: 24
  }
  value: {
    name: "CompletionItemKind_25"
    number: 25
  }
}
enum_type: {
  name: "DiagnosticTag"
  value: {
    name: "DiagnosticTag_1"
    number: 1
  }
  value: {
    name: "DiagnosticTag_2"
    number: 2
  }
}
enum_type: {
  name: "DiagnosticSeverity"
  value: {
    name: "DiagnosticSeverity_1"
    number: 1
  }
  value: {
    name: "DiagnosticSeverity_2"
    number: 2
  }
  value: {
    name: "DiagnosticSeverity_3"
    number: 3
  }
  value: {
    name: "DiagnosticSeverity_4"
    number: 4
  }
}
options: {
  java_package: "dev.lsp.go"
  java_outer_classname: "TextDocument"
  java_multiple_files: true
  go_package: "go.lsp.dev.textDocument;textDocument"
  cc_enable_arenas: true
  csharp_namespace: "Go.Lsp.Dev.TextDocument"
}
syntax: "proto3"

syntax = "proto3";

package go.lsp.dev.textDocument;

option java_package = "dev.lsp.go";

option java_outer_classname = "TextDocument";

option java_multiple_files = true;

option go_package = "go.lsp.dev.textDocument;textDocument";

option cc_enable_arenas = true;

option csharp_namespace = "Go.Lsp.Dev.TextDocument";

message BaseSymbolInformation {
  string container_name = 1;

  string name = 2;
}

message Declaration {
  oneof declaration {
    Declaration_1 declaration_1 = 1;

    Declaration_2 declaration_2 = 2;
  }

  message Declaration_1 {
    DocumentUri uri = 1;

    Range range = 2;
  }

  message Declaration_2 {
    repeated Location location = 1;

    message Location {
      Range range = 1;

      DocumentUri uri = 2;
    }
  }
}

message ChangeAnnotations {
  int32 _counter = 1;

  int32 _size = 2;

  int32 size = 3;
}

message DocumentUri {
  string document_uri = 1;
}

message Position {
  Uinteger character = 1;

  Uinteger line = 2;
}

message ChangeAnnotation {
  string description = 1;

  string label = 2;

  bool needs_confirmation = 3;
}

message ChangeAnnotationIdentifier {
  string change_annotation_identifier = 1;
}

message Uinteger {
  int32 uinteger = 1;
}

message Location {
  Range range = 1;

  DocumentUri uri = 2;
}

message Range {
  Position start = 1;

  Position end = 2;
}

message Decimal {
  int32 decimal = 1;
}

message Integer {
  int32 integer = 1;
}

message AnnotatedTextEdit {
  ChangeAnnotationIdentifier annotation_id = 1;

  string new_text = 2;

  Range range = 3;
}

enum SemanticTokenTypes {
  SemanticTokenTypes_Class = 1;

  SemanticTokenTypes_Comment = 2;

  SemanticTokenTypes_Decorator = 3;

  SemanticTokenTypes_Enum = 4;

  SemanticTokenTypes_EnumMember = 5;

  SemanticTokenTypes_Event = 6;

  SemanticTokenTypes_Function = 7;

  SemanticTokenTypes_Interface = 8;

  SemanticTokenTypes_Keyword = 9;

  SemanticTokenTypes_Macro = 10;

  SemanticTokenTypes_Method = 11;

  SemanticTokenTypes_Modifier = 12;

  SemanticTokenTypes_Namespace = 13;

  SemanticTokenTypes_Number = 14;

  SemanticTokenTypes_Operator = 15;

  SemanticTokenTypes_Parameter = 16;

  SemanticTokenTypes_Property = 17;

  SemanticTokenTypes_Regexp = 18;

  SemanticTokenTypes_String = 19;

  SemanticTokenTypes_Struct = 20;

  SemanticTokenTypes_Type = 21;

  SemanticTokenTypes_TypeParameter = 22;

  SemanticTokenTypes_Variable = 23;
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

enum Tags {
  Tags_1 = 1;
}

enum InlayHintKind {
  InlayHintKind_1 = 1;

  InlayHintKind_2 = 2;
}

enum SemanticTokenModifiers {
  SemanticTokenModifiers_Abstract = 1;

  SemanticTokenModifiers_Async = 2;

  SemanticTokenModifiers_Declaration = 3;

  SemanticTokenModifiers_DefaultLibrary = 4;

  SemanticTokenModifiers_Definition = 5;

  SemanticTokenModifiers_Deprecated = 6;

  SemanticTokenModifiers_Documentation = 7;

  SemanticTokenModifiers_Modification = 8;

  SemanticTokenModifiers_Readonly = 9;

  SemanticTokenModifiers_Static = 10;
}

enum DocumentHighlightKind {
  DocumentHighlightKind_1 = 1;

  DocumentHighlightKind_2 = 2;

  DocumentHighlightKind_3 = 3;
}

enum CompletionItemKind {
  CompletionItemKind_1 = 1;

  CompletionItemKind_2 = 2;

  CompletionItemKind_3 = 3;

  CompletionItemKind_4 = 4;

  CompletionItemKind_5 = 5;

  CompletionItemKind_6 = 6;

  CompletionItemKind_7 = 7;

  CompletionItemKind_8 = 8;

  CompletionItemKind_9 = 9;

  CompletionItemKind_10 = 10;

  CompletionItemKind_11 = 11;

  CompletionItemKind_12 = 12;

  CompletionItemKind_13 = 13;

  CompletionItemKind_14 = 14;

  CompletionItemKind_15 = 15;

  CompletionItemKind_16 = 16;

  CompletionItemKind_17 = 17;

  CompletionItemKind_18 = 18;

  CompletionItemKind_19 = 19;

  CompletionItemKind_20 = 20;

  CompletionItemKind_21 = 21;

  CompletionItemKind_22 = 22;

  CompletionItemKind_23 = 23;

  CompletionItemKind_24 = 24;

  CompletionItemKind_25 = 25;
}

enum DiagnosticTag {
  DiagnosticTag_1 = 1;

  DiagnosticTag_2 = 2;
}

enum DiagnosticSeverity {
  DiagnosticSeverity_1 = 1;

  DiagnosticSeverity_2 = 2;

  DiagnosticSeverity_3 = 3;

  DiagnosticSeverity_4 = 4;
}
```

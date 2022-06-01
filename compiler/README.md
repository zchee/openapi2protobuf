```proto
syntax = "proto3";

package go.lsp.dev.textDocument;

option java_outer_classname = "TextDocument";

option java_multiple_files = true;

option go_package = "go.lsp.dev.textDocument;textDocument";

option cc_enable_arenas = true;

option csharp_namespace = "Go.Lsp.Dev.TextDocument";

option java_package = "dev.lsp.go";

message Type {
  Range range = 1;

  string text = 2;
}

message OptionalVersionedTextDocumentIdentifier {
  DocumentUri uri = 1;

  int32 version = 2;
}

message ChangeAnnotationIdentifier {
  string change_annotation_identifier = 1;
}

message CodeActionKind {
  string code_action_kind = 1;
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

message SemanticTokensEdit {
  Uinteger delete_count = 1;

  Uinteger start = 2;

  repeated int32 data = 3;
}

message Hover {
  Contents contents = 1;

  Range range = 2;

  message Contents {
    oneof contents {
      contents_1 contents_1 = 1;

      contents_2 contents_2 = 2;

      contents_3 contents_3 = 3;

      contents_4 contents_4 = 4;
    }

    message contents_1 {
      string value = 1;
    }

    message contents_2 {
      string language = 1;

      string value = 2;
    }

    message contents_3 {
      repeated   = 1;

      message  {
        oneof  {
          ._1 _1 = 1;

          ._2 _2 = 2;
        }

        message _1 {
          string language = 1;

          string value = 2;
        }

        message _2 {
          string  = 1;
        }
      }
    }

    message contents_4 {
      string  = 1;
    }
  }
}

message InlineValue {
  oneof inline_value {
    InlineValue_1 inline_value_1 = 1;

    InlineValue_2 inline_value_2 = 2;

    InlineValue_3 inline_value_3 = 3;
  }

  message InlineValue_1 {
    Range range = 1;

    string text = 2;
  }

  message InlineValue_2 {
    string variable_name = 1;

    bool case_sensitive_lookup = 2;

    Range range = 3;
  }

  message InlineValue_3 {
    string expression = 1;

    Range range = 2;
  }
}

message InsertReplaceEdit {
  Range insert = 1;

  string new_text = 2;

  Range replace = 3;
}

message CompletionItemLabelDetails {
  string description = 1;

  string detail = 2;
}

message FormattingOptions {
  bool insert_final_newline = 1;

  bool insert_spaces = 2;

  Uinteger tab_size = 3;

  bool trim_final_newlines = 4;

  bool trim_trailing_whitespace = 5;
}

message MarkupContent {
  string value = 1;
}

message TextDocumentEdit {
  repeated Edits edits = 1;

  OptionalVersionedTextDocumentIdentifier text_document = 2;

  message Edits {
    repeated  edits = 1;

    message Edits {
      oneof edits {
        Edits_1 edits_1 = 1;

        Edits_2 edits_2 = 2;
      }

      message Edits_1 {
        string new_text = 1;

        Range range = 2;
      }

      message Edits_2 {
        ChangeAnnotationIdentifier annotation_id = 1;

        string new_text = 2;

        Range range = 3;
      }
    }
  }
}

message WorkspaceFolder {
  URI uri = 1;

  string name = 2;
}

message CreateFile {
  ChangeAnnotationIdentifier annotation_id = 1;

  CreateFileOptions options = 2;

  DocumentUri uri = 3;
}

message Type1 {
  Range range = 1;

  string variable_name = 2;

  bool case_sensitive_lookup = 3;
}

message ChangeAnnotation {
  string description = 1;

  string label = 2;

  bool needs_confirmation = 3;
}

message FullTextDocument {
  repeated int32 _line_offsets = 1;

  string language_id = 2;

  Integer version = 3;

  string _content = 4;

  string _language_id = 5;

  DocumentUri _uri = 6;

  Integer _version = 7;

  int32 line_count = 8;

  string uri = 9;
}

message SemanticTokensLegend {
  repeated string token_modifiers = 1;

  repeated string token_types = 2;
}

message RenameFile {
  ChangeAnnotationIdentifier annotation_id = 1;

  DocumentUri new_uri = 2;

  DocumentUri old_uri = 3;

  RenameFileOptions options = 4;
}

message SymbolInformation {
  Location location = 1;

  string name = 2;

  string container_name = 3;

  bool deprecated = 4;
}

message TextDocumentContentChangeEvent {
  oneof text_document_content_change_event {
    TextDocumentContentChangeEvent_1 text_document_content_change_event_1 = 1;

    TextDocumentContentChangeEvent_2 text_document_content_change_event_2 = 2;
  }

  message TextDocumentContentChangeEvent_1 {
    Range range = 1;

    Uinteger range_length = 2;

    string text = 3;
  }

  message TextDocumentContentChangeEvent_2 {
    string text = 1;
  }
}

message WorkspaceChange {
  ChangeAnnotations _change_annotations = 1;

  WorkspaceEdit _workspace_edit = 2;

  WorkspaceEdit edit = 3;
}

message Command {
  string command = 1;

  string title = 2;
}

message Definition {
  oneof definition {
    Definition_1 definition_1 = 1;

    Definition_2 definition_2 = 2;
  }

  message Definition_1 {
    Range range = 1;

    DocumentUri uri = 2;
  }

  message Definition_2 {
    repeated Location location = 1;

    message Location {
      Range range = 1;

      DocumentUri uri = 2;
    }
  }
}

message ReferenceContext {
  bool include_declaration = 1;
}

message BaseSymbolInformation {
  string container_name = 1;

  string name = 2;
}

message DocumentHighlight {
  Range range = 1;
}

message Position {
  Uinteger character = 1;

  Uinteger line = 2;
}

message TextDocumentIdentifier {
  DocumentUri uri = 1;
}

message TextEdit {
  string new_text = 1;

  Range range = 2;
}

message URI {
  string uri = 1;
}

message Integer {
  int32 integer = 1;
}

message AnnotatedTextEdit {
  ChangeAnnotationIdentifier annotation_id = 1;

  string new_text = 2;

  Range range = 3;
}

message ResourceOperation {
  ChangeAnnotationIdentifier annotation_id = 1;

  string kind = 2;
}

message ColorPresentation {
  string label = 1;

  TextEdit text_edit = 2;

  repeated AdditionalTextEdits additional_text_edits = 3;

  message AdditionalTextEdits {
    repeated TextEdit text_edit = 1;

    message TextEdit {
      string new_text = 1;

      Range range = 2;
    }
  }
}

message DeleteFileOptions {
  bool ignore_if_not_exists = 1;

  bool recursive = 2;
}

message DocumentUri {
  string document_uri = 1;
}

message SemanticTokens {
  repeated int32 data = 1;

  string result_id = 2;
}

message Color {
  Decimal alpha = 1;

  Decimal blue = 2;

  Decimal green = 3;

  Decimal red = 4;
}

message CreateFileOptions {
  bool ignore_if_exists = 1;

  bool overwrite = 2;
}

message LocationLink {
  Range origin_selection_range = 1;

  Range target_range = 2;

  Range target_selection_range = 3;

  DocumentUri target_uri = 4;
}

message  {
  Command command = 1;

  Location location = 2;

  .Tooltip tooltip = 3;

  string value = 4;

  message Tooltip {
    oneof tooltip {
      tooltip_1 tooltip_1 = 1;

      tooltip_2 tooltip_2 = 2;
    }

    message tooltip_1 {
      string value = 1;
    }

    message tooltip_2 {
      string  = 1;
    }
  }
}

message TextDocument {
  string language_id = 1;

  Uinteger line_count = 2;

  DocumentUri uri = 3;

  Integer version = 4;
}

message DiagnosticRelatedInformation {
  Location location = 1;

  string message = 2;
}

message FoldingRangeKind {
  string folding_range_kind = 1;
}

message Decimal {
  int32 decimal = 1;
}

message VersionedTextDocumentIdentifier {
  Integer version = 1;

  DocumentUri uri = 2;
}

message Location {
  Range range = 1;

  DocumentUri uri = 2;
}

message ChangeAnnotations {
  int32 _counter = 1;

  int32 _size = 2;

  int32 size = 3;
}

message ColorInformation {
  Range range = 1;

  Color color = 2;
}

message DeleteFile {
  ChangeAnnotationIdentifier annotation_id = 1;

  DeleteFileOptions options = 2;

  DocumentUri uri = 3;
}

message RenameFileOptions {
  bool overwrite = 1;

  bool ignore_if_exists = 2;
}

message Type2 {
  string expression = 1;

  Range range = 2;
}

message CodeDescription {
  URI href = 1;
}

message MarkedString {
  oneof marked_string {
    MarkedString_1 marked_string_1 = 1;

    MarkedString_2 marked_string_2 = 2;
  }

  message MarkedString_1 {
    string language = 1;

    string value = 2;
  }

  message MarkedString_2 {
    string  = 1;
  }
}

message Range {
  Position end = 1;

  Position start = 2;
}

message SemanticTokensDelta {
  repeated Edits edits = 1;

  string result_id = 2;

  message Edits {
    repeated SemanticTokensEdit semantic_tokens_edit = 1;

    message SemanticTokensEdit {
      repeated int32 data = 1;

      Uinteger delete_count = 2;

      Uinteger start = 3;
    }
  }
}

message TextDocumentItem {
  string language_id = 1;

  string text = 2;

  DocumentUri uri = 3;

  Integer version = 4;
}

message TextEditChangeImpl {
  ChangeAnnotations change_annotations = 1;

  repeated Edits edits = 2;

  message Edits {
    repeated  edits = 1;

    message Edits {
      oneof edits {
        Edits_1 edits_1 = 1;

        Edits_2 edits_2 = 2;
      }

      message Edits_1 {
        string new_text = 1;

        Range range = 2;
      }

      message Edits_2 {
        string new_text = 1;

        Range range = 2;

        ChangeAnnotationIdentifier annotation_id = 3;
      }
    }
  }
}

message WorkspaceEdit {
  repeated DocumentChanges document_changes = 1;

  message DocumentChanges {
    repeated  document_changes = 1;

    message DocumentChanges {
      oneof document_changes {
        DocumentChanges_1 document_changes_1 = 1;

        DocumentChanges_2 document_changes_2 = 2;

        DocumentChanges_3 document_changes_3 = 3;

        DocumentChanges_4 document_changes_4 = 4;
      }

      message DocumentChanges_1 {
        repeated Edits edits = 1;

        OptionalVersionedTextDocumentIdentifier text_document = 2;

        message Edits {
          repeated  edits = 1;

          message Edits {
            oneof edits {
              Edits_1 edits_1 = 1;

              Edits_2 edits_2 = 2;
            }

            message Edits_1 {
              string new_text = 1;

              Range range = 2;
            }

            message Edits_2 {
              ChangeAnnotationIdentifier annotation_id = 1;

              string new_text = 2;

              Range range = 3;
            }
          }
        }
      }

      message DocumentChanges_2 {
        ChangeAnnotationIdentifier annotation_id = 1;

        CreateFileOptions options = 2;

        DocumentUri uri = 3;
      }

      message DocumentChanges_3 {
        DocumentUri new_uri = 1;

        DocumentUri old_uri = 2;

        RenameFileOptions options = 3;

        ChangeAnnotationIdentifier annotation_id = 4;
      }

      message DocumentChanges_4 {
        DocumentUri uri = 1;

        ChangeAnnotationIdentifier annotation_id = 2;

        DeleteFileOptions options = 3;
      }
    }
  }
}

message FoldingRange {
  string collapsed_text = 1;

  Uinteger end_character = 2;

  Uinteger end_line = 3;

  FoldingRangeKind kind = 4;

  Uinteger start_character = 5;

  Uinteger start_line = 6;
}

message Uinteger {
  int32 uinteger = 1;
}

enum MarkupKind {
  MarkupKind_Markdown = 1;

  MarkupKind_Plaintext = 2;
}

enum InlayHintKind {
  InlayHintKind_1 = 1;

  InlayHintKind_2 = 2;
}

enum DiagnosticTag {
  DiagnosticTag_1 = 1;

  DiagnosticTag_2 = 2;
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

enum Kind {
  Kind_Create = 1;
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

enum DiagnosticSeverity {
  DiagnosticSeverity_1 = 1;

  DiagnosticSeverity_2 = 2;

  DiagnosticSeverity_3 = 3;

  DiagnosticSeverity_4 = 4;
}

enum Tags {
  Tags_1 = 1;
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

enum SymbolTag {
  SymbolTag_1 = 1;
}

enum DocumentHighlightKind {
  DocumentHighlightKind_1 = 1;

  DocumentHighlightKind_2 = 2;

  DocumentHighlightKind_3 = 3;
}

enum CompletionItemTag {
  CompletionItemTag_1 = 1;
}

enum InsertTextMode {
  InsertTextMode_1 = 1;

  InsertTextMode_2 = 2;
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

enum CodeActionTriggerKind {
  CodeActionTriggerKind_1 = 1;

  CodeActionTriggerKind_2 = 2;
}

enum InsertTextFormat {
  InsertTextFormat_1 = 1;

  InsertTextFormat_2 = 2;
}
```

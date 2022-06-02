```proto
syntax = "proto3";

package go.lsp.dev.types;

import "google/protobuf/any.proto";

option java_package = "dev.lsp.go";

option java_outer_classname = "Types";

option java_multiple_files = true;

option go_package = "go.lsp.dev.types";

option cc_enable_arenas = true;

option csharp_namespace = "Go.Lsp.Dev.Types";

message  {
}

message AnnotatedTextEdit {
  ChangeAnnotationIdentifier annotation_id = 1;

  string new_text = 2;

  Range range = 3;
}

message BaseSymbolInformation {
  string container_name = 1;

  string name = 2;

  repeated Tags tags = 3;

  message Tags {
  }
}

message CallHierarchyIncomingCall {
  CallHierarchyItem from = 1;

  repeated FromRanges from_ranges = 2;

  message FromRanges {
    repeated Range range = 1;

    message Range {
      Position end = 1;

      Position start = 2;
    }
  }
}

message CallHierarchyItem {
  DocumentUri uri = 1;

  string detail = 2;

  string name = 3;

  Range range = 4;

  Range selection_range = 5;

  repeated Tags tags = 6;

  message Tags {
  }
}

message CallHierarchyOutgoingCall {
  repeated FromRanges from_ranges = 1;

  CallHierarchyItem to = 2;

  message FromRanges {
    repeated Range range = 1;

    message Range {
      Position end = 1;

      Position start = 2;
    }
  }
}

message ChangeAnnotation {
  string description = 1;

  string label = 2;

  bool needs_confirmation = 3;
}

message ChangeAnnotationIdentifier {
  string change_annotation_identifier = 1;
}

message ChangeAnnotations {
  Annotations _annotations = 1;

  int32 _counter = 2;

  int32 _size = 3;

  int32 size = 4;

  message Annotations {
  }
}

message CodeActionKind {
  string code_action_kind = 1;
}

message CodeDescription {
  URI href = 1;
}

message CodeLens {
  Range range = 1;

  Command command = 2;
}

message Color {
  Decimal blue = 1;

  Decimal green = 2;

  Decimal red = 3;

  Decimal alpha = 4;
}

message ColorInformation {
  Color color = 1;

  Range range = 2;
}

message ColorPresentation {
  TextEdit text_edit = 1;

  repeated AdditionalTextEdits additional_text_edits = 2;

  string label = 3;

  message AdditionalTextEdits {
    repeated TextEdit text_edit = 1;

    message TextEdit {
      string new_text = 1;

      Range range = 2;
    }
  }
}

message Command {
  repeated Arguments arguments = 1;

  string command = 2;

  string title = 3;

  message Arguments {
  }
}

message CompletionItem {
  repeated string commit_characters = 1;

  Documentation documentation = 2;

  CompletionItemLabelDetails label_details = 3;

  repeated Tags tags = 4;

  repeated AdditionalTextEdits additional_text_edits = 5;

  bool deprecated = 6;

  string detail = 7;

  string insert_text = 8;

  bool preselect = 9;

  Command command = 10;

  string filter_text = 11;

  string label = 12;

  string sort_text = 13;

  TextEdit text_edit = 14;

  string text_edit_text = 15;

  message Documentation {
    oneof documentation {
      MarkupContent markup_content = 1;

      documentation_2 documentation_2 = 2;
    }

    message MarkupContent {
      string value = 1;
    }

    message documentation_2 {
      string  = 1;
    }
  }

  message Tags {
  }

  message AdditionalTextEdits {
    repeated TextEdit text_edit = 1;

    message TextEdit {
      string new_text = 1;

      Range range = 2;
    }
  }

  message TextEdit {
    oneof text_edit {
      TextEdit text_edit = 1;

      InsertReplaceEdit insert_replace_edit = 2;
    }

    message TextEdit {
      string new_text = 1;

      Range range = 2;
    }

    message InsertReplaceEdit {
      Range insert = 1;

      string new_text = 2;

      Range replace = 3;
    }
  }
}

message CompletionItemLabelDetails {
  string description = 1;

  string detail = 2;
}

message CompletionList {
  ItemDefaults item_defaults = 1;

  repeated Items items = 2;

  bool is_incomplete = 3;

  message ItemDefaults {
    repeated string commit_characters = 1;

    EditRange edit_range = 2;

    message EditRange {
      oneof edit_range {
        Range range = 1;

        editRange_2 edit_range_2 = 2;
      }

      message Range {
        Position end = 1;

        Position start = 2;
      }

      message editRange_2 {
        Range insert = 1;

        Range replace = 2;
      }
    }
  }

  message Items {
    repeated CompletionItem completion_item = 1;

    message CompletionItem {
      string insert_text = 1;

      bool preselect = 2;

      TextEdit text_edit = 3;

      string text_edit_text = 4;

      Command command = 5;

      string filter_text = 6;

      string label = 7;

      string sort_text = 8;

      repeated string commit_characters = 9;

      Documentation documentation = 10;

      CompletionItemLabelDetails label_details = 11;

      repeated Tags tags = 12;

      repeated AdditionalTextEdits additional_text_edits = 13;

      bool deprecated = 14;

      string detail = 15;

      message TextEdit {
        oneof text_edit {
          TextEdit text_edit = 1;

          InsertReplaceEdit insert_replace_edit = 2;
        }

        message TextEdit {
          string new_text = 1;

          Range range = 2;
        }

        message InsertReplaceEdit {
          string new_text = 1;

          Range replace = 2;

          Range insert = 3;
        }
      }

      message Documentation {
        oneof documentation {
          MarkupContent markup_content = 1;

          documentation_2 documentation_2 = 2;
        }

        message MarkupContent {
          string value = 1;
        }

        message documentation_2 {
          string  = 1;
        }
      }

      message Tags {
      }

      message AdditionalTextEdits {
        repeated TextEdit text_edit = 1;

        message TextEdit {
          string new_text = 1;

          Range range = 2;
        }
      }
    }
  }
}

message CreateFile {
  ChangeAnnotationIdentifier annotation_id = 1;

  CreateFileOptions options = 2;

  DocumentUri uri = 3;
}

message CreateFileOptions {
  bool overwrite = 1;

  bool ignore_if_exists = 2;
}

message Decimal {
  int32 decimal = 1;
}

message Declaration {
  oneof declaration {
    Location location = 1;

    Declaration_2 declaration_2 = 2;
  }

  message Location {
    Range range = 1;

    DocumentUri uri = 2;
  }

  message Declaration_2 {
    repeated Location location = 1;

    message Location {
      DocumentUri uri = 1;

      Range range = 2;
    }
  }
}

message Definition {
  oneof definition {
    Location location = 1;

    Definition_2 definition_2 = 2;
  }

  message Location {
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

message DeleteFile {
  DeleteFileOptions options = 1;

  DocumentUri uri = 2;

  ChangeAnnotationIdentifier annotation_id = 3;
}

message DeleteFileOptions {
  bool ignore_if_not_exists = 1;

  bool recursive = 2;
}

message DiagnosticRelatedInformation {
  Location location = 1;

  string message = 2;
}

message DocumentHighlight {
  Range range = 1;
}

message DocumentLink {
  Range range = 1;

  string target = 2;

  string tooltip = 3;
}

message DocumentUri {
  string document_uri = 1;
}

message FoldingRange {
  string collapsed_text = 1;

  Uinteger end_character = 2;

  Uinteger end_line = 3;

  FoldingRangeKind kind = 4;

  Uinteger start_character = 5;

  Uinteger start_line = 6;
}

message FoldingRangeKind {
  string folding_range_kind = 1;
}

message FormattingOptions {
  bool insert_final_newline = 1;

  bool insert_spaces = 2;

  Uinteger tab_size = 3;

  bool trim_final_newlines = 4;

  bool trim_trailing_whitespace = 5;
}

message FullTextDocument {
  string _content = 1;

  DocumentUri _uri = 2;

  int32 line_count = 3;

  string uri = 4;

  Integer version = 5;

  string _language_id = 6;

  repeated int32 _line_offsets = 7;

  Integer _version = 8;

  string language_id = 9;
}

message Hover {
  Contents contents = 1;

  Range range = 2;

  message Contents {
    oneof contents {
      MarkupContent markup_content = 1;

      contents_2 contents_2 = 2;

      contents_3 contents_3 = 3;

      contents_4 contents_4 = 4;
    }

    message MarkupContent {
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
    InlineValueText inline_value_text = 1;

    InlineValueVariableLookup inline_value_variable_lookup = 2;

    InlineValueEvaluatableExpression inline_value_evaluatable_expression = 3;
  }

  message InlineValueText {
    Range range = 1;

    string text = 2;
  }

  message InlineValueVariableLookup {
    bool case_sensitive_lookup = 1;

    Range range = 2;

    string variable_name = 3;
  }

  message InlineValueEvaluatableExpression {
    Range range = 1;

    string expression = 2;
  }
}

message InlineValueEvaluatableExpression {
  string expression = 1;

  Range range = 2;
}

message InlineValueText {
  Range range = 1;

  string text = 2;
}

message InlineValueVariableLookup {
  bool case_sensitive_lookup = 1;

  Range range = 2;

  string variable_name = 3;
}

message InsertReplaceEdit {
  Range replace = 1;

  Range insert = 2;

  string new_text = 3;
}

message Integer {
  int32 integer = 1;
}

message LSPAny {
  google.protobuf.Any any = 1;
}

message LSPArray {
  repeated google.protobuf.Any array = 1;
}

message LSPObject {
  option map_entry = true;

  map<string, google.protobuf.Any> object = 1;
}

message Location {
  DocumentUri uri = 1;

  Range range = 2;
}

message LocationLink {
  Range origin_selection_range = 1;

  Range target_range = 2;

  Range target_selection_range = 3;

  DocumentUri target_uri = 4;
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

message MarkupContent {
  string value = 1;
}

message OptionalVersionedTextDocumentIdentifier {
  DocumentUri uri = 1;

  int32 version = 2;
}

message Position {
  Uinteger character = 1;

  Uinteger line = 2;
}

message Range {
  Position end = 1;

  Position start = 2;
}

message ReferenceContext {
  bool include_declaration = 1;
}

message RenameFile {
  ChangeAnnotationIdentifier annotation_id = 1;

  DocumentUri new_uri = 2;

  DocumentUri old_uri = 3;

  RenameFileOptions options = 4;
}

message RenameFileOptions {
  bool ignore_if_exists = 1;

  bool overwrite = 2;
}

message ResourceOperation {
  ChangeAnnotationIdentifier annotation_id = 1;

  string kind = 2;
}

message SemanticTokens {
  repeated int32 data = 1;

  string result_id = 2;
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

message SemanticTokensEdit {
  repeated int32 data = 1;

  Uinteger delete_count = 2;

  Uinteger start = 3;
}

message SemanticTokensLegend {
  repeated string token_modifiers = 1;

  repeated string token_types = 2;
}

message SymbolInformation {
  Location location = 1;

  string name = 2;

  repeated Tags tags = 3;

  string container_name = 4;

  bool deprecated = 5;

  message Tags {
  }
}

message TextDocument {
  Integer version = 1;

  string language_id = 2;

  Uinteger line_count = 3;

  DocumentUri uri = 4;
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

message TextDocumentEdit {
  repeated Edits edits = 1;

  OptionalVersionedTextDocumentIdentifier text_document = 2;

  message Edits {
    repeated  edits = 1;

    message Edits {
      oneof edits {
        TextEdit text_edit = 1;

        AnnotatedTextEdit annotated_text_edit = 2;
      }

      message TextEdit {
        string new_text = 1;

        Range range = 2;
      }

      message AnnotatedTextEdit {
        ChangeAnnotationIdentifier annotation_id = 1;

        string new_text = 2;

        Range range = 3;
      }
    }
  }
}

message TextDocumentIdentifier {
  DocumentUri uri = 1;
}

message TextDocumentItem {
  string language_id = 1;

  string text = 2;

  DocumentUri uri = 3;

  Integer version = 4;
}

message TextEdit {
  string new_text = 1;

  Range range = 2;
}

message TextEditChange {
}

message TextEditChangeImpl {
  ChangeAnnotations change_annotations = 1;

  repeated Edits edits = 2;

  message Edits {
    repeated  edits = 1;

    message Edits {
      oneof edits {
        TextEdit text_edit = 1;

        AnnotatedTextEdit annotated_text_edit = 2;
      }

      message TextEdit {
        string new_text = 1;

        Range range = 2;
      }

      message AnnotatedTextEdit {
        ChangeAnnotationIdentifier annotation_id = 1;

        string new_text = 2;

        Range range = 3;
      }
    }
  }
}

message URI {
  string uri = 1;
}

message Uinteger {
  int32 uinteger = 1;
}

message VersionedTextDocumentIdentifier {
  DocumentUri uri = 1;

  Integer version = 2;
}

message WorkspaceChange {
  WorkspaceEdit edit = 1;

  ChangeAnnotations _change_annotations = 2;

  TextEditChanges _text_edit_changes = 3;

  WorkspaceEdit _workspace_edit = 4;

  message TextEditChanges {
  }
}

message WorkspaceEdit {
  Changes changes = 1;

  repeated DocumentChanges document_changes = 2;

  ChangeAnnotations change_annotations = 3;

  message Changes {
  }

  message DocumentChanges {
    repeated  document_changes = 1;

    message DocumentChanges {
      oneof document_changes {
        TextDocumentEdit text_document_edit = 1;

        CreateFile create_file = 2;

        RenameFile rename_file = 3;

        DeleteFile delete_file = 4;
      }

      message TextDocumentEdit {
        repeated Edits edits = 1;

        OptionalVersionedTextDocumentIdentifier text_document = 2;

        message Edits {
          repeated  edits = 1;

          message Edits {
            oneof edits {
              TextEdit text_edit = 1;

              AnnotatedTextEdit annotated_text_edit = 2;
            }

            message TextEdit {
              string new_text = 1;

              Range range = 2;
            }

            message AnnotatedTextEdit {
              ChangeAnnotationIdentifier annotation_id = 1;

              string new_text = 2;

              Range range = 3;
            }
          }
        }
      }

      message CreateFile {
        CreateFileOptions options = 1;

        DocumentUri uri = 2;

        ChangeAnnotationIdentifier annotation_id = 3;
      }

      message RenameFile {
        ChangeAnnotationIdentifier annotation_id = 1;

        DocumentUri new_uri = 2;

        DocumentUri old_uri = 3;

        RenameFileOptions options = 4;
      }

      message DeleteFile {
        DocumentUri uri = 1;

        ChangeAnnotationIdentifier annotation_id = 2;

        DeleteFileOptions options = 3;
      }
    }
  }

  message ChangeAnnotations {
  }
}

message WorkspaceFolder {
  string name = 1;

  URI uri = 2;
}

message WorkspaceSymbol {
  string container_name = 1;

  Location location = 2;

  string name = 3;

  repeated Tags tags = 4;

  message Location {
    oneof location {
      Location location = 1;

      location_2 location_2 = 2;
    }

    message Location {
      Range range = 1;

      DocumentUri uri = 2;
    }

    message location_2 {
      DocumentUri uri = 1;
    }
  }

  message Tags {
  }
}

enum CodeActionTriggerKind {
  CodeActionTriggerKind_1 = 1;

  CodeActionTriggerKind_2 = 2;
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

enum CompletionItemTag {
  CompletionItemTag_1 = 1;
}

enum DiagnosticSeverity {
  DiagnosticSeverity_1 = 1;

  DiagnosticSeverity_2 = 2;

  DiagnosticSeverity_3 = 3;

  DiagnosticSeverity_4 = 4;
}

enum DiagnosticTag {
  DiagnosticTag_1 = 1;

  DiagnosticTag_2 = 2;
}

enum DocumentHighlightKind {
  DocumentHighlightKind_1 = 1;

  DocumentHighlightKind_2 = 2;

  DocumentHighlightKind_3 = 3;
}

enum InlayHintKind {
  InlayHintKind_1 = 1;

  InlayHintKind_2 = 2;
}

enum InsertTextFormat {
  InsertTextFormat_1 = 1;

  InsertTextFormat_2 = 2;
}

enum InsertTextMode {
  InsertTextMode_1 = 1;

  InsertTextMode_2 = 2;
}

enum Kind {
  Kind_Rename = 1;
}

enum MarkupKind {
  MarkupKind_Markdown = 1;

  MarkupKind_Plaintext = 2;
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

enum SymbolTag {
  SymbolTag_1 = 1;
}

enum Tags {
  Tags_1 = 1;
}
```

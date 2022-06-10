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

message AnnotatedTextEdit {
  ChangeAnnotationIdentifier annotation_id = 1;

  Range range = 2;

  string new_text = 3;
}

message BaseSymbolInformation {
  string name = 1;

  SymbolKind kind = 2;

  repeated Tags tags = 3;

  string container_name = 4;

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;
      }
    }
  }
}

message CallHierarchyIncomingCall {
  CallHierarchyItem from = 1;

  repeated FromRanges from_ranges = 2;

  message FromRanges {
    Range range = 1;
  }
}

message CallHierarchyItem {
  string name = 1;

  SymbolKind kind = 2;

  repeated Tags tags = 3;

  string detail = 4;

  DocumentUri uri = 5;

  Range range = 6;

  Range selection_range = 7;

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;
      }
    }
  }
}

message CallHierarchyOutgoingCall {
  CallHierarchyItem to = 1;

  repeated FromRanges from_ranges = 2;

  message FromRanges {
    Range range = 1;
  }
}

message ChangeAnnotation {
  string label = 1;

  bool needs_confirmation = 2;

  string description = 3;
}

message ChangeAnnotationIdentifier {
  string change_annotation_identifier = 1;
}

message ChangeAnnotations {
  ChangeAnnotation _annotations = 1;

  int32 _counter = 2;

  int32 _size = 3;

  int32 size = 4;

  message ChangeAnnotation {
    string label = 1;

    bool needs_confirmation = 2;

    string description = 3;
  }
}

message CodeAction {
  string title = 1;

  CodeActionKind kind = 2;

  repeated Diagnostics diagnostics = 3;

  bool is_preferred = 4;

  string disabled = 5;

  WorkspaceEdit edit = 6;

  Command command = 7;

  message Diagnostics {
    Diagnostic diagnostic = 1;
  }
}

message CodeActionContext {
  repeated Diagnostics diagnostics = 1;

  repeated Only only = 2;

  CodeActionTriggerKind trigger_kind = 3;

  message Diagnostics {
    Diagnostic diagnostic = 1;
  }

  message Only {
    Only only = 1;

    message Only {
      string only = 1;
    }
  }
}

message CodeActionKind {
  string code_action_kind = 1;
}

message CodeActionTriggerKind {
  enum CodeActionTriggerKind {
    CodeActionTriggerKind_1 = 1;

    CodeActionTriggerKind_2 = 2;
  }
}

message CodeDescription {
  URI href = 1;
}

message CodeLens {
  Range range = 1;

  Command command = 2;
}

message Color {
  Decimal red = 1;

  Decimal green = 2;

  Decimal blue = 3;

  Decimal alpha = 4;
}

message ColorInformation {
  Range range = 1;

  Color color = 2;
}

message ColorPresentation {
  string label = 1;

  TextEdit text_edit = 2;

  repeated AdditionalTextEdits additional_text_edits = 3;

  message AdditionalTextEdits {
    TextEdit text_edit = 1;
  }
}

message Command {
  string title = 1;

  string command = 2;

  repeated Arguments arguments = 3;

  message Arguments {
    LSPAny lsp_any = 1;
  }
}

message CompletionItem {
  string label = 1;

  CompletionItemLabelDetails label_details = 2;

  CompletionItemKind kind = 3;

  repeated Tags tags = 4;

  string detail = 5;

  Documentation documentation = 6;

  bool deprecated = 7;

  bool preselect = 8;

  string sort_text = 9;

  string filter_text = 10;

  string insert_text = 11;

  InsertTextFormat insert_text_format = 12;

  InsertTextMode insert_text_mode = 13;

  TextEdit text_edit = 14;

  string text_edit_text = 15;

  repeated AdditionalTextEdits additional_text_edits = 16;

  repeated CommitCharacters commit_characters = 17;

  Command command = 18;

  message AdditionalTextEdits {
    TextEdit text_edit = 1;
  }

  message Documentation {
    oneof documentation {
      MarkupContent markup_content = 1;

      Documentation2 documentation_2 = 2;
    }

    message Documentation2 {
      string documentation_2 = 1;
    }
  }

  message TextEdit {
    oneof text_edit {
      TextEdit text_edit = 1;

      InsertReplaceEdit insert_replace_edit = 2;
    }
  }

  message CommitCharacters {
    CommitCharacters commit_characters = 1;

    message CommitCharacters {
      string commit_characters = 1;
    }
  }

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;
      }
    }
  }
}

message CompletionItemKind {
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
}

message CompletionItemLabelDetails {
  string detail = 1;

  string description = 2;
}

message CompletionItemTag {
  enum CompletionItemTag {
    CompletionItemTag_1 = 1;
  }
}

message CompletionList {
  bool is_incomplete = 1;

  ItemDefaults item_defaults = 2;

  repeated Items items = 3;

  message ItemDefaults {
    EditRange edit_range = 1;

    InsertTextFormat insert_text_format = 2;

    InsertTextMode insert_text_mode = 3;

    repeated CommitCharacters commit_characters = 4;

    message EditRange {
      oneof edit_range {
        Range range = 1;

        EditRange2 edit_range_2 = 2;
      }

      message EditRange2 {
        Range insert = 1;

        Range replace = 2;
      }
    }

    message CommitCharacters {
      CommitCharacters commit_characters = 1;

      message CommitCharacters {
        string commit_characters = 1;
      }
    }
  }

  message Items {
    CompletionItem completion_item = 1;
  }
}

message CreateFile {
  Kind kind = 1;

  DocumentUri uri = 2;

  CreateFileOptions options = 3;

  ChangeAnnotationIdentifier annotation_id = 4;

  message Kind {
    enum Kind {
      Kind_Create = 1;
    }
  }
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

    Declaration2 declaration_2 = 2;
  }

  message Declaration2 {
    Location location = 1;
  }
}

message Definition {
  oneof definition {
    Location location = 1;

    Definition2 definition_2 = 2;
  }

  message Definition2 {
    Location location = 1;
  }
}

message DeleteFile {
  Kind kind = 1;

  DocumentUri uri = 2;

  DeleteFileOptions options = 3;

  ChangeAnnotationIdentifier annotation_id = 4;

  message Kind {
    enum Kind {
      Kind_Delete = 1;
    }
  }
}

message DeleteFileOptions {
  bool recursive = 1;

  bool ignore_if_not_exists = 2;
}

message Diagnostic {
  Range range = 1;

  DiagnosticSeverity severity = 2;

  Code code = 3;

  CodeDescription code_description = 4;

  string source = 5;

  string message = 6;

  repeated Tags tags = 7;

  repeated RelatedInformation related_information = 8;

  message RelatedInformation {
    DiagnosticRelatedInformation diagnostic_related_information = 1;
  }

  message Code {
    oneof code {
      Code1 code_1 = 1;

      Code2 code_2 = 2;
    }

    message Code1 {
      string code_1 = 1;
    }

    message Code2 {
      int32 code_2 = 1;
    }
  }

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;

        Tags_2 = 2;
      }
    }
  }
}

message DiagnosticRelatedInformation {
  Location location = 1;

  string message = 2;
}

message DiagnosticSeverity {
  enum DiagnosticSeverity {
    DiagnosticSeverity_1 = 1;

    DiagnosticSeverity_2 = 2;

    DiagnosticSeverity_3 = 3;

    DiagnosticSeverity_4 = 4;
  }
}

message DiagnosticTag {
  enum DiagnosticTag {
    DiagnosticTag_1 = 1;

    DiagnosticTag_2 = 2;
  }
}

message DocumentHighlight {
  Range range = 1;

  DocumentHighlightKind kind = 2;
}

message DocumentHighlightKind {
  enum DocumentHighlightKind {
    DocumentHighlightKind_1 = 1;

    DocumentHighlightKind_2 = 2;

    DocumentHighlightKind_3 = 3;
  }
}

message DocumentLink {
  Range range = 1;

  string target = 2;

  string tooltip = 3;
}

message DocumentSymbol {
  string name = 1;

  string detail = 2;

  SymbolKind kind = 3;

  repeated Tags tags = 4;

  bool deprecated = 5;

  Range range = 6;

  Range selection_range = 7;

  repeated Children children = 8;

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;
      }
    }
  }

  message Children {
    DocumentSymbol document_symbol = 1;
  }
}

message DocumentUri {
  string document_uri = 1;
}

message FoldingRange {
  Uinteger start_line = 1;

  Uinteger start_character = 2;

  Uinteger end_line = 3;

  Uinteger end_character = 4;

  FoldingRangeKind kind = 5;

  string collapsed_text = 6;
}

message FoldingRangeKind {
  string folding_range_kind = 1;
}

message FormattingOptions {
  Uinteger tab_size = 1;

  bool insert_spaces = 2;

  bool trim_trailing_whitespace = 3;

  bool insert_final_newline = 4;

  bool trim_final_newlines = 5;
}

message FullTextDocument {
  DocumentUri _uri = 1;

  string _language_id = 2;

  Integer _version = 3;

  string _content = 4;

  repeated LineOffsets _line_offsets = 5;

  string uri = 6;

  string language_id = 7;

  Integer version = 8;

  int32 line_count = 9;

  message LineOffsets {
    LineOffsets line_offsets = 1;

    message LineOffsets {
      int32 line_offsets = 1;
    }
  }
}

message Hover {
  Contents contents = 1;

  Range range = 2;

  message Contents {
    oneof contents {
      MarkupContent markup_content = 1;

      Contents2 contents_2 = 2;

      Contents3 contents_3 = 3;

      Contents4 contents_4 = 4;
    }

    message Contents2 {
      string value = 1;

      string language = 2;
    }

    message Contents3 {
      Contents3 contents_3 = 1;

      message Contents3 {
        oneof contents_3 {
          Contents31 contents_31 = 1;

          Contents32 contents_32 = 2;
        }

        message Contents31 {
          string language = 1;

          string value = 2;
        }

        message Contents32 {
          string contents_3_2 = 1;
        }
      }
    }

    message Contents4 {
      string contents_4 = 1;
    }
  }
}

message InlayHint {
  Position position = 1;

  Label label = 2;

  InlayHintKind kind = 3;

  repeated TextEdits text_edits = 4;

  Tooltip tooltip = 5;

  bool padding_left = 6;

  bool padding_right = 7;

  message TextEdits {
    TextEdit text_edit = 1;
  }

  message Tooltip {
    oneof tooltip {
      MarkupContent markup_content = 1;

      Tooltip2 tooltip_2 = 2;
    }

    message Tooltip2 {
      string tooltip_2 = 1;
    }
  }

  message Label {
    oneof label {
      Label1 label_1 = 1;

      Label2 label_2 = 2;
    }

    message Label1 {
      Label1 label_1 = 1;

      message Label1 {
        Location location = 1;

        Tooltip tooltip = 2;

        string value = 3;

        Command command = 4;

        message Tooltip {
          oneof tooltip {
            MarkupContent markup_content = 1;

            Tooltip2 tooltip_2 = 2;
          }

          message Tooltip2 {
            string tooltip_2 = 1;
          }
        }
      }
    }

    message Label2 {
      string label_2 = 1;
    }
  }
}

message InlayHintKind {
  enum InlayHintKind {
    InlayHintKind_1 = 1;

    InlayHintKind_2 = 2;
  }
}

message InlayHintLabelPart {
  string value = 1;

  Tooltip tooltip = 2;

  Location location = 3;

  Command command = 4;

  message Tooltip {
    oneof tooltip {
      MarkupContent markup_content = 1;

      Tooltip2 tooltip_2 = 2;
    }

    message Tooltip2 {
      string tooltip_2 = 1;
    }
  }
}

message InlineValue {
  oneof inline_value {
    InlineValueText inline_value_text = 1;

    InlineValueVariableLookup inline_value_variable_lookup = 2;

    InlineValueEvaluatableExpression inline_value_evaluatable_expression = 3;
  }
}

message InlineValueContext {
  int32 frame_id = 1;

  Range stopped_location = 2;
}

message InlineValueEvaluatableExpression {
  Range range = 1;

  string expression = 2;
}

message InlineValueText {
  Range range = 1;

  string text = 2;
}

message InlineValueVariableLookup {
  Range range = 1;

  string variable_name = 2;

  bool case_sensitive_lookup = 3;
}

message InsertReplaceEdit {
  string new_text = 1;

  Range insert = 2;

  Range replace = 3;
}

message InsertTextFormat {
  enum InsertTextFormat {
    InsertTextFormat_1 = 1;

    InsertTextFormat_2 = 2;
  }
}

message InsertTextMode {
  enum InsertTextMode {
    InsertTextMode_1 = 1;

    InsertTextMode_2 = 2;
  }
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

  DocumentUri target_uri = 2;

  Range target_range = 3;

  Range target_selection_range = 4;
}

message MarkedString {
  oneof marked_string {
    MarkedString1 marked_string_1 = 1;

    MarkedString2 marked_string_2 = 2;
  }

  message MarkedString1 {
    string language = 1;

    string value = 2;
  }

  message MarkedString2 {
    string marked_string_2 = 1;
  }
}

message MarkupContent {
  MarkupKind kind = 1;

  string value = 2;
}

message MarkupKind {
  enum MarkupKind {
    MarkupKind_Markdown = 1;

    MarkupKind_Plaintext = 2;
  }
}

message OptionalVersionedTextDocumentIdentifier {
  int32 version = 1;

  DocumentUri uri = 2;
}

message Position {
  Uinteger line = 1;

  Uinteger character = 2;
}

message Range {
  Position start = 1;

  Position end = 2;
}

message ReferenceContext {
  bool include_declaration = 1;
}

message RenameFile {
  Kind kind = 1;

  DocumentUri old_uri = 2;

  DocumentUri new_uri = 3;

  RenameFileOptions options = 4;

  ChangeAnnotationIdentifier annotation_id = 5;

  message Kind {
    enum Kind {
      Kind_Rename = 1;
    }
  }
}

message RenameFileOptions {
  bool overwrite = 1;

  bool ignore_if_exists = 2;
}

message ResourceOperation {
  string kind = 1;

  ChangeAnnotationIdentifier annotation_id = 2;
}

message SemanticTokenModifiers {
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
}

message SemanticTokenTypes {
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
}

message SemanticTokens {
  string result_id = 1;

  repeated Data data = 2;

  message Data {
    Data data = 1;

    message Data {
      int32 data = 1;
    }
  }
}

message SemanticTokensDelta {
  string result_id = 1;

  repeated Edits edits = 2;

  message Edits {
    SemanticTokensEdit semantic_tokens_edit = 1;
  }
}

message SemanticTokensEdit {
  Uinteger start = 1;

  Uinteger delete_count = 2;

  repeated Data data = 3;

  message Data {
    Data data = 1;

    message Data {
      int32 data = 1;
    }
  }
}

message SemanticTokensLegend {
  repeated TokenTypes token_types = 1;

  repeated TokenModifiers token_modifiers = 2;

  message TokenTypes {
    TokenTypes token_types = 1;

    message TokenTypes {
      string token_types = 1;
    }
  }

  message TokenModifiers {
    TokenModifiers token_modifiers = 1;

    message TokenModifiers {
      string token_modifiers = 1;
    }
  }
}

message SymbolInformation {
  bool deprecated = 1;

  Location location = 2;

  string name = 3;

  SymbolKind kind = 4;

  repeated Tags tags = 5;

  string container_name = 6;

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;
      }
    }
  }
}

message SymbolKind {
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
}

message SymbolTag {
  enum SymbolTag {
    SymbolTag_1 = 1;
  }
}

message TextDocument {
  DocumentUri uri = 1;

  string language_id = 2;

  Integer version = 3;

  Uinteger line_count = 4;
}

message TextDocumentContentChangeEvent {
  oneof text_document_content_change_event {
    TextDocumentContentChangeEvent1 text_document_content_change_event_1 = 1;

    TextDocumentContentChangeEvent2 text_document_content_change_event_2 = 2;
  }

  message TextDocumentContentChangeEvent1 {
    Uinteger range_length = 1;

    string text = 2;

    Range range = 3;
  }

  message TextDocumentContentChangeEvent2 {
    string text = 1;
  }
}

message TextDocumentEdit {
  OptionalVersionedTextDocumentIdentifier text_document = 1;

  repeated Edits edits = 2;

  message Edits {
    Edits edits = 1;

    message Edits {
      oneof edits {
        TextEdit text_edit = 1;

        AnnotatedTextEdit annotated_text_edit = 2;
      }
    }
  }
}

message TextDocumentIdentifier {
  DocumentUri uri = 1;
}

message TextDocumentItem {
  DocumentUri uri = 1;

  string language_id = 2;

  Integer version = 3;

  string text = 4;
}

message TextEdit {
  Range range = 1;

  string new_text = 2;
}

message TextEditChangeImpl {
  repeated Edits edits = 1;

  ChangeAnnotations change_annotations = 2;

  message Edits {
    Edits edits = 1;

    message Edits {
      oneof edits {
        TextEdit text_edit = 1;

        AnnotatedTextEdit annotated_text_edit = 2;
      }
    }
  }
}

message TypeHierarchyItem {
  string name = 1;

  SymbolKind kind = 2;

  repeated Tags tags = 3;

  string detail = 4;

  DocumentUri uri = 5;

  Range range = 6;

  Range selection_range = 7;

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;
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
  Integer version = 1;

  DocumentUri uri = 2;
}

message WorkspaceChange {
  WorkspaceEdit _workspace_edit = 1;

  TextEditChangeImpl _text_edit_changes = 2;

  ChangeAnnotations _change_annotations = 3;

  WorkspaceEdit edit = 4;

  message TextEditChangeImpl {
    ChangeAnnotations change_annotations = 1;

    repeated Edits edits = 2;

    message Edits {
      Edits edits = 1;

      message Edits {
        oneof edits {
          TextEdit text_edit = 1;

          AnnotatedTextEdit annotated_text_edit = 2;
        }
      }
    }
  }
}

message WorkspaceEdit {
  AdditionalProperties changes = 1;

  repeated DocumentChanges document_changes = 2;

  ChangeAnnotation change_annotations = 3;

  message ChangeAnnotation {
    string description = 1;

    string label = 2;

    bool needs_confirmation = 3;
  }

  message AdditionalProperties {
    TextEdit text_edit = 1;
  }

  message DocumentChanges {
    DocumentChanges document_changes = 1;

    message DocumentChanges {
      oneof document_changes {
        TextDocumentEdit text_document_edit = 1;

        CreateFile create_file = 2;

        RenameFile rename_file = 3;

        DeleteFile delete_file = 4;
      }
    }
  }
}

message WorkspaceFolder {
  URI uri = 1;

  string name = 2;
}

message WorkspaceSymbol {
  Location location = 1;

  string name = 2;

  SymbolKind kind = 3;

  repeated Tags tags = 4;

  string container_name = 5;

  message Location {
    oneof location {
      Location location = 1;

      Location2 location_2 = 2;
    }

    message Location2 {
      DocumentUri uri = 1;
    }
  }

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;
      }
    }
  }
}
```

```proto
syntax = "proto3";

package go.lsp.dev.types;

import "google/protobuf/any.proto";

option csharp_namespace = "Go.Lsp.Dev.Types";

option java_package = "dev.lsp.go";

option java_outer_classname = "Types";

option java_multiple_files = true;

option go_package = "go.lsp.dev.types";

option cc_enable_arenas = true;

message AnnotatedTextEdit {
  ChangeAnnotationIdentifier annotation_id = 1;

  Range range = 2;

  string new_text = 3;
}

message ApplyWorkspaceEditParams {
  string label = 1;

  WorkspaceEdit edit = 2;
}

message ApplyWorkspaceEditResult {
  bool applied = 1;

  string failure_reason = 2;

  Uinteger failed_change = 3;
}

message CallHierarchyClientCapabilities {
  bool dynamic_registration = 1;
}

message CallHierarchyOptions {
  bool work_done_progress = 1;
}

message CallHierarchyRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  string id = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message ChangeAnnotation {
  string label = 1;

  bool needs_confirmation = 2;

  string description = 3;
}

message ChangeAnnotationIdentifier {
  string change_annotation_identifier = 1;
}

message ClientCapabilities {
  WorkspaceClientCapabilities workspace = 1;

  TextDocumentClientCapabilities text_document = 2;

  NotebookDocumentClientCapabilities notebook_document = 3;

  WindowClientCapabilities window = 4;

  GeneralClientCapabilities general = 5;
}

message CodeActionClientCapabilities {
  bool dynamic_registration = 1;

  CodeActionLiteralSupport code_action_literal_support = 2;

  bool is_preferred_support = 3;

  bool disabled_support = 4;

  bool data_support = 5;

  ResolveSupport resolve_support = 6;

  bool honors_change_annotations = 7;

  message ResolveSupport {
    repeated Properties properties = 1;

    message Properties {
      Properties properties = 1;

      message Properties {
        string properties = 1;
      }
    }
  }

  message CodeActionLiteralSupport {
    CodeActionKind code_action_kind = 1;

    message CodeActionKind {
      repeated ValueSet value_set = 1;

      message ValueSet {
        ValueSet value_set = 1;

        message ValueSet {
          string value_set = 1;
        }
      }
    }
  }
}

message CodeActionContext {
  repeated Diagnostics diagnostics = 1;

  repeated Only only = 2;

  CodeActionTriggerKind trigger_kind = 3;

  message Diagnostics {
    Diagnostic diagnostic = 1;
  }

  message Only {
    Only only = 1;

    message Only {
      string only = 1;
    }
  }
}

message CodeActionOptions {
  repeated CodeActionKinds code_action_kinds = 1;

  bool resolve_provider = 2;

  bool work_done_progress = 3;

  message CodeActionKinds {
    CodeActionKinds code_action_kinds = 1;

    message CodeActionKinds {
      string code_action_kinds = 1;
    }
  }
}

message CodeActionParams {
  TextDocumentIdentifier text_document = 1;

  Range range = 2;

  CodeActionContext context = 3;

  ProgressToken work_done_token = 4;

  ProgressToken partial_result_token = 5;
}

message CodeActionRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  repeated CodeActionKinds code_action_kinds = 2;

  bool resolve_provider = 3;

  bool work_done_progress = 4;

  message CodeActionKinds {
    CodeActionKinds code_action_kinds = 1;

    message CodeActionKinds {
      string code_action_kinds = 1;
    }
  }

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message CodeActionTriggerKind {
  enum CodeActionTriggerKind {
    CodeActionTriggerKind_1 = 1;

    CodeActionTriggerKind_2 = 2;
  }
}

message CodeDescription {
  URI href = 1;
}

message CodeLensClientCapabilities {
  bool dynamic_registration = 1;
}

message CodeLensOptions {
  bool resolve_provider = 1;

  bool work_done_progress = 2;
}

message CodeLensParams {
  TextDocumentIdentifier text_document = 1;

  ProgressToken work_done_token = 2;

  ProgressToken partial_result_token = 3;
}

message CodeLensRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool resolve_provider = 2;

  bool work_done_progress = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message CodeLensWorkspaceClientCapabilities {
  bool refresh_support = 1;
}

message CompletionClientCapabilities {
  bool dynamic_registration = 1;

  CompletionItem completion_item = 2;

  CompletionItemKind completion_item_kind = 3;

  InsertTextMode insert_text_mode = 4;

  bool context_support = 5;

  CompletionList completion_list = 6;

  message CompletionItem {
    bool deprecated_support = 1;

    bool preselect_support = 2;

    bool snippet_support = 3;

    bool commit_characters_support = 4;

    bool insert_replace_support = 5;

    InsertTextModeSupport insert_text_mode_support = 6;

    bool label_details_support = 7;

    ResolveSupport resolve_support = 8;

    TagSupport tag_support = 9;

    repeated DocumentationFormat documentation_format = 10;

    message InsertTextModeSupport {
      repeated ValueSet value_set = 1;

      message ValueSet {
        ValueSet value_set = 1;

        message ValueSet {
          enum ValueSet {
            ValueSet_1 = 1;

            ValueSet_2 = 2;
          }
        }
      }
    }

    message ResolveSupport {
      repeated Properties properties = 1;

      message Properties {
        Properties properties = 1;

        message Properties {
          string properties = 1;
        }
      }
    }

    message TagSupport {
      repeated ValueSet value_set = 1;

      message ValueSet {
        ValueSet value_set = 1;

        message ValueSet {
          enum ValueSet {
            ValueSet_1 = 1;
          }
        }
      }
    }

    message DocumentationFormat {
      DocumentationFormat documentation_format = 1;

      message DocumentationFormat {
        enum DocumentationFormat {
          DocumentationFormat_Markdown = 1;

          DocumentationFormat_Plaintext = 2;
        }
      }
    }
  }

  message CompletionItemKind {
    repeated ValueSet value_set = 1;

    message ValueSet {
      ValueSet value_set = 1;

      message ValueSet {
        enum ValueSet {
          ValueSet_1 = 1;

          ValueSet_2 = 2;

          ValueSet_3 = 3;

          ValueSet_4 = 4;

          ValueSet_5 = 5;

          ValueSet_6 = 6;

          ValueSet_7 = 7;

          ValueSet_8 = 8;

          ValueSet_9 = 9;

          ValueSet_10 = 10;

          ValueSet_11 = 11;

          ValueSet_12 = 12;

          ValueSet_13 = 13;

          ValueSet_14 = 14;

          ValueSet_15 = 15;

          ValueSet_16 = 16;

          ValueSet_17 = 17;

          ValueSet_18 = 18;

          ValueSet_19 = 19;

          ValueSet_20 = 20;

          ValueSet_21 = 21;

          ValueSet_22 = 22;

          ValueSet_23 = 23;

          ValueSet_24 = 24;

          ValueSet_25 = 25;
        }
      }
    }
  }

  message CompletionList {
    repeated ItemDefaults item_defaults = 1;

    message ItemDefaults {
      ItemDefaults item_defaults = 1;

      message ItemDefaults {
        string item_defaults = 1;
      }
    }
  }
}

message CompletionContext {
  CompletionTriggerKind trigger_kind = 1;

  string trigger_character = 2;
}

message CompletionOptions {
  repeated TriggerCharacters trigger_characters = 1;

  repeated AllCommitCharacters all_commit_characters = 2;

  bool resolve_provider = 3;

  bool completion_item = 4;

  bool work_done_progress = 5;

  message AllCommitCharacters {
    AllCommitCharacters all_commit_characters = 1;

    message AllCommitCharacters {
      string all_commit_characters = 1;
    }
  }

  message TriggerCharacters {
    TriggerCharacters trigger_characters = 1;

    message TriggerCharacters {
      string trigger_characters = 1;
    }
  }
}

message CompletionParams {
  CompletionContext context = 1;

  TextDocumentIdentifier text_document = 2;

  Position position = 3;

  ProgressToken work_done_token = 4;

  ProgressToken partial_result_token = 5;
}

message CompletionRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  repeated TriggerCharacters trigger_characters = 2;

  repeated AllCommitCharacters all_commit_characters = 3;

  bool resolve_provider = 4;

  bool completion_item = 5;

  bool work_done_progress = 6;

  message TriggerCharacters {
    TriggerCharacters trigger_characters = 1;

    message TriggerCharacters {
      string trigger_characters = 1;
    }
  }

  message AllCommitCharacters {
    AllCommitCharacters all_commit_characters = 1;

    message AllCommitCharacters {
      string all_commit_characters = 1;
    }
  }

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message CompletionTriggerKind {
  enum CompletionTriggerKind {
    CompletionTriggerKind_1 = 1;

    CompletionTriggerKind_2 = 2;

    CompletionTriggerKind_3 = 3;
  }
}

message CreateFile {
  Kind kind = 1;

  DocumentUri uri = 2;

  CreateFileOptions options = 3;

  ChangeAnnotationIdentifier annotation_id = 4;

  message Kind {
    enum Kind {
      Kind_Create = 1;
    }
  }
}

message CreateFileOptions {
  bool overwrite = 1;

  bool ignore_if_exists = 2;
}

message DeclarationClientCapabilities {
  bool dynamic_registration = 1;

  bool link_support = 2;
}

message DeclarationOptions {
  bool work_done_progress = 1;
}

message DeclarationRegistrationOptions {
  bool work_done_progress = 1;

  repeated DocumentSelector document_selector = 2;

  string id = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message DefinitionClientCapabilities {
  bool dynamic_registration = 1;

  bool link_support = 2;
}

message DefinitionOptions {
  bool work_done_progress = 1;
}

message DefinitionParams {
  TextDocumentIdentifier text_document = 1;

  Position position = 2;

  ProgressToken work_done_token = 3;

  ProgressToken partial_result_token = 4;
}

message DefinitionRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message DeleteFile {
  Kind kind = 1;

  DocumentUri uri = 2;

  DeleteFileOptions options = 3;

  ChangeAnnotationIdentifier annotation_id = 4;

  message Kind {
    enum Kind {
      Kind_Delete = 1;
    }
  }
}

message DeleteFileOptions {
  bool recursive = 1;

  bool ignore_if_not_exists = 2;
}

message Diagnostic {
  Range range = 1;

  DiagnosticSeverity severity = 2;

  Code code = 3;

  CodeDescription code_description = 4;

  string source = 5;

  string message = 6;

  repeated Tags tags = 7;

  repeated RelatedInformation related_information = 8;

  message Tags {
    Tags tags = 1;

    message Tags {
      enum Tags {
        Tags_1 = 1;

        Tags_2 = 2;
      }
    }
  }

  message Code {
    oneof code {
      Code1 code_1 = 1;

      Code2 code_2 = 2;
    }

    message Code1 {
      string code_1 = 1;
    }

    message Code2 {
      int32 code_2 = 1;
    }
  }

  message RelatedInformation {
    DiagnosticRelatedInformation diagnostic_related_information = 1;
  }
}

message DiagnosticClientCapabilities {
  Notebook notebook = 1;

  string language = 2;

  message Notebook {
    oneof notebook {
      Notebook1 notebook_1 = 1;

      Notebook2 notebook_2 = 2;

      Notebook3 notebook_3 = 3;

      Notebook4 notebook_4 = 4;
    }

    message Notebook1 {
      string notebook_type = 1;

      string pattern = 2;

      string scheme = 3;
    }

    message Notebook2 {
      string notebook_type = 1;

      string pattern = 2;

      string scheme = 3;
    }

    message Notebook3 {
      string notebook_type = 1;

      string pattern = 2;

      string scheme = 3;
    }

    message Notebook4 {
      string notebook_4 = 1;
    }
  }
}

message DiagnosticRelatedInformation {
  Location location = 1;

  string message = 2;
}

message DiagnosticSeverity {
  enum DiagnosticSeverity {
    DiagnosticSeverity_1 = 1;

    DiagnosticSeverity_2 = 2;

    DiagnosticSeverity_3 = 3;

    DiagnosticSeverity_4 = 4;
  }
}

message DiagnosticWorkspaceClientCapabilities {
  bool refresh_support = 1;
}

message DidChangeConfigurationClientCapabilities {
  bool dynamic_registration = 1;
}

message DidChangeConfigurationRegistrationOptions {
  Section section = 1;

  message Section {
    oneof section {
      Section1 section_1 = 1;

      Section2 section_2 = 2;
    }

    message Section1 {
      Section1 section_1 = 1;

      message Section1 {
        string section_1 = 1;
      }
    }

    message Section2 {
      string section_2 = 1;
    }
  }
}

message DidChangeTextDocumentParams {
  VersionedTextDocumentIdentifier text_document = 1;

  repeated ContentChanges content_changes = 2;

  message ContentChanges {
    ContentChanges content_changes = 1;

    message ContentChanges {
      oneof content_changes {
        ContentChanges1 content_changes_1 = 1;

        ContentChanges2 content_changes_2 = 2;
      }

      message ContentChanges1 {
        Range range = 1;

        Uinteger range_length = 2;

        string text = 3;
      }

      message ContentChanges2 {
        string text = 1;
      }
    }
  }
}

message DidChangeWatchedFilesClientCapabilities {
  bool dynamic_registration = 1;

  bool relative_pattern_support = 2;
}

message DidChangeWatchedFilesParams {
  repeated Changes changes = 1;

  message Changes {
    FileEvent file_event = 1;
  }
}

message DidChangeWatchedFilesRegistrationOptions {
  repeated Watchers watchers = 1;

  message Watchers {
    FileSystemWatcher file_system_watcher = 1;
  }
}

message DidCloseTextDocumentParams {
  TextDocumentIdentifier text_document = 1;
}

message DidOpenTextDocumentParams {
  TextDocumentItem text_document = 1;
}

message DidSaveTextDocumentParams {
  TextDocumentIdentifier text_document = 1;

  string text = 2;
}

message DocumentColorClientCapabilities {
  bool dynamic_registration = 1;
}

message DocumentColorOptions {
  bool work_done_progress = 1;
}

message DocumentColorRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  string id = 2;

  bool work_done_progress = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message DocumentFilter {
  oneof document_filter {
    DocumentFilter1 document_filter_1 = 1;

    DocumentFilter2 document_filter_2 = 2;

    DocumentFilter3 document_filter_3 = 3;

    NotebookCellTextDocumentFilter notebook_cell_text_document_filter = 4;
  }

  message DocumentFilter1 {
    string language = 1;

    string pattern = 2;

    string scheme = 3;
  }

  message DocumentFilter2 {
    string scheme = 1;

    string language = 2;

    string pattern = 3;
  }

  message DocumentFilter3 {
    string scheme = 1;

    string language = 2;

    string pattern = 3;
  }
}

message DocumentFormattingClientCapabilities {
  bool dynamic_registration = 1;
}

message DocumentFormattingOptions {
  bool work_done_progress = 1;
}

message DocumentFormattingParams {
  TextDocumentIdentifier text_document = 1;

  FormattingOptions options = 2;

  ProgressToken work_done_token = 3;
}

message DocumentFormattingRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message DocumentHighlightClientCapabilities {
  bool dynamic_registration = 1;
}

message DocumentHighlightOptions {
  bool work_done_progress = 1;
}

message DocumentHighlightParams {
  TextDocumentIdentifier text_document = 1;

  Position position = 2;

  ProgressToken work_done_token = 3;

  ProgressToken partial_result_token = 4;
}

message DocumentHighlightRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message DocumentLinkClientCapabilities {
  bool dynamic_registration = 1;

  bool tooltip_support = 2;
}

message DocumentLinkOptions {
  bool resolve_provider = 1;

  bool work_done_progress = 2;
}

message DocumentLinkParams {
  TextDocumentIdentifier text_document = 1;

  ProgressToken work_done_token = 2;

  ProgressToken partial_result_token = 3;
}

message DocumentLinkRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool resolve_provider = 2;

  bool work_done_progress = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message DocumentOnTypeFormattingClientCapabilities {
  bool dynamic_registration = 1;
}

message DocumentOnTypeFormattingOptions {
  string first_trigger_character = 1;

  repeated MoreTriggerCharacter more_trigger_character = 2;

  message MoreTriggerCharacter {
    MoreTriggerCharacter more_trigger_character = 1;

    message MoreTriggerCharacter {
      string more_trigger_character = 1;
    }
  }
}

message DocumentOnTypeFormattingParams {
  TextDocumentIdentifier text_document = 1;

  Position position = 2;

  string ch = 3;

  FormattingOptions options = 4;
}

message DocumentOnTypeFormattingRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  string first_trigger_character = 2;

  repeated MoreTriggerCharacter more_trigger_character = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }

  message MoreTriggerCharacter {
    MoreTriggerCharacter more_trigger_character = 1;

    message MoreTriggerCharacter {
      string more_trigger_character = 1;
    }
  }
}

message DocumentRangeFormattingClientCapabilities {
  bool dynamic_registration = 1;
}

message DocumentRangeFormattingOptions {
  bool work_done_progress = 1;
}

message DocumentRangeFormattingParams {
  TextDocumentIdentifier text_document = 1;

  Range range = 2;

  FormattingOptions options = 3;

  ProgressToken work_done_token = 4;
}

message DocumentRangeFormattingRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message DocumentSelector {
  DocumentSelector document_selector = 1;
}

message DocumentSymbolClientCapabilities {
  bool dynamic_registration = 1;

  SymbolKind symbol_kind = 2;

  bool hierarchical_document_symbol_support = 3;

  TagSupport tag_support = 4;

  bool label_support = 5;

  message SymbolKind {
    repeated ValueSet value_set = 1;

    message ValueSet {
      ValueSet value_set = 1;

      message ValueSet {
        enum ValueSet {
          ValueSet_1 = 1;

          ValueSet_2 = 2;

          ValueSet_3 = 3;

          ValueSet_4 = 4;

          ValueSet_5 = 5;

          ValueSet_6 = 6;

          ValueSet_7 = 7;

          ValueSet_8 = 8;

          ValueSet_9 = 9;

          ValueSet_10 = 10;

          ValueSet_11 = 11;

          ValueSet_12 = 12;

          ValueSet_13 = 13;

          ValueSet_14 = 14;

          ValueSet_15 = 15;

          ValueSet_16 = 16;

          ValueSet_17 = 17;

          ValueSet_18 = 18;

          ValueSet_19 = 19;

          ValueSet_20 = 20;

          ValueSet_21 = 21;

          ValueSet_22 = 22;

          ValueSet_23 = 23;

          ValueSet_24 = 24;

          ValueSet_25 = 25;

          ValueSet_26 = 26;
        }
      }
    }
  }

  message TagSupport {
    repeated ValueSet value_set = 1;

    message ValueSet {
      ValueSet value_set = 1;

      message ValueSet {
        enum ValueSet {
          ValueSet_1 = 1;
        }
      }
    }
  }
}

message DocumentSymbolOptions {
  string label = 1;

  bool work_done_progress = 2;
}

message DocumentSymbolParams {
  TextDocumentIdentifier text_document = 1;

  ProgressToken work_done_token = 2;

  ProgressToken partial_result_token = 3;
}

message DocumentSymbolRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  string label = 2;

  bool work_done_progress = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message DocumentUri {
  string document_uri = 1;
}

message ExecuteCommandClientCapabilities {
  bool dynamic_registration = 1;
}

message ExecuteCommandOptions {
  repeated Commands commands = 1;

  bool work_done_progress = 2;

  message Commands {
    Commands commands = 1;

    message Commands {
      string commands = 1;
    }
  }
}

message ExecuteCommandParams {
  string command = 1;

  ProgressToken work_done_token = 2;
}

message ExecuteCommandRegistrationOptions {
  repeated Commands commands = 1;

  bool work_done_progress = 2;

  message Commands {
    Commands commands = 1;

    message Commands {
      string commands = 1;
    }
  }
}

message FailureHandlingKind {
  enum FailureHandlingKind {
    FailureHandlingKind_Abort = 1;

    FailureHandlingKind_TextOnlyTransactional = 2;

    FailureHandlingKind_Transactional = 3;

    FailureHandlingKind_Undo = 4;
  }
}

message FileChangeType {
  enum FileChangeType {
    FileChangeType_1 = 1;

    FileChangeType_2 = 2;

    FileChangeType_3 = 3;
  }
}

message FileEvent {
  DocumentUri uri = 1;

  FileChangeType type = 2;
}

message FileOperationClientCapabilities {
  bool dynamic_registration = 1;

  bool did_create = 2;

  bool will_create = 3;

  bool did_rename = 4;

  bool will_rename = 5;

  bool did_delete = 6;

  bool will_delete = 7;
}

message FileOperationFilter {
  string scheme = 1;

  FileOperationPattern pattern = 2;
}

message FileOperationOptions {
  FileOperationRegistrationOptions did_create = 1;

  FileOperationRegistrationOptions will_create = 2;

  FileOperationRegistrationOptions did_rename = 3;

  FileOperationRegistrationOptions will_rename = 4;

  FileOperationRegistrationOptions did_delete = 5;

  FileOperationRegistrationOptions will_delete = 6;
}

message FileOperationPattern {
  string glob = 1;

  FileOperationPatternKind matches = 2;

  FileOperationPatternOptions options = 3;
}

message FileOperationPatternKind {
  enum FileOperationPatternKind {
    FileOperationPatternKind_File = 1;

    FileOperationPatternKind_Folder = 2;
  }
}

message FileOperationPatternOptions {
  bool ignore_case = 1;
}

message FileOperationRegistrationOptions {
  repeated Filters filters = 1;

  message Filters {
    FileOperationFilter file_operation_filter = 1;
  }
}

message FileSystemWatcher {
  GlobPattern glob_pattern = 1;

  WatchKind kind = 2;
}

message FoldingRangeClientCapabilities {
  bool dynamic_registration = 1;

  Uinteger range_limit = 2;

  bool line_folding_only = 3;

  FoldingRangeKind folding_range_kind = 4;

  bool folding_range = 5;

  message FoldingRangeKind {
    repeated ValueSet value_set = 1;

    message ValueSet {
      ValueSet value_set = 1;

      message ValueSet {
        string value_set = 1;
      }
    }
  }
}

message FoldingRangeOptions {
  bool work_done_progress = 1;
}

message FoldingRangeRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  string id = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message FormattingOptions {
  Uinteger tab_size = 1;

  bool insert_spaces = 2;

  bool trim_trailing_whitespace = 3;

  bool insert_final_newline = 4;

  bool trim_final_newlines = 5;
}

message GeneralClientCapabilities {
  StaleRequestSupport stale_request_support = 1;

  RegularExpressionsClientCapabilities regular_expressions = 2;

  MarkdownClientCapabilities markdown = 3;

  repeated PositionEncodings position_encodings = 4;

  message PositionEncodings {
    PositionEncodings position_encodings = 1;

    message PositionEncodings {
      string position_encodings = 1;
    }
  }

  message StaleRequestSupport {
    bool cancel = 1;

    repeated RetryOnContentModified retry_on_content_modified = 2;

    message RetryOnContentModified {
      RetryOnContentModified retry_on_content_modified = 1;

      message RetryOnContentModified {
        string retry_on_content_modified = 1;
      }
    }
  }
}

message GlobPattern {
  oneof glob_pattern {
    RelativePattern relative_pattern = 1;

    GlobPattern2 glob_pattern_2 = 2;
  }

  message GlobPattern2 {
    string glob_pattern_2 = 1;
  }
}

message HoverClientCapabilities {
  bool dynamic_registration = 1;

  repeated ContentFormat content_format = 2;

  message ContentFormat {
    ContentFormat content_format = 1;

    message ContentFormat {
      enum ContentFormat {
        ContentFormat_Markdown = 1;

        ContentFormat_Plaintext = 2;
      }
    }
  }
}

message HoverOptions {
  bool work_done_progress = 1;
}

message HoverParams {
  TextDocumentIdentifier text_document = 1;

  Position position = 2;

  ProgressToken work_done_token = 3;
}

message HoverRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message ImplementationClientCapabilities {
  bool dynamic_registration = 1;

  bool link_support = 2;
}

message ImplementationOptions {
  bool work_done_progress = 1;
}

message ImplementationRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  string id = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message InitializeError {
  bool retry = 1;
}

message InitializeErrorCodes {
  enum InitializeErrorCodes {
    InitializeErrorCodes_1 = 1;
  }
}

message InitializeParams {
  int32 process_id = 1;

  ClientInfo client_info = 2;

  string locale = 3;

  string root_path = 4;

  string root_uri = 5;

  ClientCapabilities capabilities = 6;

  Trace trace = 7;

  ProgressToken work_done_token = 8;

  message Trace {
    enum Trace {
      Trace_Compact = 1;

      Trace_Messages = 2;

      Trace_Off = 3;

      Trace_Verbose = 4;
    }
  }

  message ClientInfo {
    string name = 1;

    string version = 2;
  }
}

message InlayHintClientCapabilities {
  bool dynamic_registration = 1;

  ResolveSupport resolve_support = 2;

  message ResolveSupport {
    repeated Properties properties = 1;

    message Properties {
      Properties properties = 1;

      message Properties {
        string properties = 1;
      }
    }
  }
}

message InlayHintWorkspaceClientCapabilities {
  bool refresh_support = 1;
}

message InlineValueClientCapabilities {
  bool dynamic_registration = 1;
}

message InlineValueWorkspaceClientCapabilities {
  bool refresh_support = 1;
}

message InsertTextMode {
  enum InsertTextMode {
    InsertTextMode_1 = 1;

    InsertTextMode_2 = 2;
  }
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

message LinkedEditingRangeClientCapabilities {
  bool dynamic_registration = 1;
}

message LinkedEditingRangeOptions {
  bool work_done_progress = 1;
}

message LinkedEditingRangeRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  string id = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message Location {
  DocumentUri uri = 1;

  Range range = 2;
}

message LogMessageParams {
  MessageType type = 1;

  string message = 2;
}

message MarkdownClientCapabilities {
  string parser = 1;

  string version = 2;

  repeated AllowedTags allowed_tags = 3;

  message AllowedTags {
    AllowedTags allowed_tags = 1;

    message AllowedTags {
      string allowed_tags = 1;
    }
  }
}

message MarkupContent {
  MarkupKind kind = 1;

  string value = 2;
}

message MarkupKind {
  enum MarkupKind {
    MarkupKind_Markdown = 1;

    MarkupKind_Plaintext = 2;
  }
}

message MessageActionItem {
  string title = 1;
}

message MessageType {
  enum MessageType {
    MessageType_1 = 1;

    MessageType_2 = 2;

    MessageType_3 = 3;

    MessageType_4 = 4;
  }
}

message MonikerClientCapabilities {
  bool dynamic_registration = 1;
}

message MonikerOptions {
  bool work_done_progress = 1;
}

message MonikerRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message NotebookCellTextDocumentFilter {
  string language = 1;

  Notebook notebook = 2;

  message Notebook {
    oneof notebook {
      Notebook1 notebook_1 = 1;

      Notebook2 notebook_2 = 2;

      Notebook3 notebook_3 = 3;

      Notebook4 notebook_4 = 4;
    }

    message Notebook1 {
      string scheme = 1;

      string notebook_type = 2;

      string pattern = 3;
    }

    message Notebook2 {
      string scheme = 1;

      string notebook_type = 2;

      string pattern = 3;
    }

    message Notebook3 {
      string notebook_type = 1;

      string pattern = 2;

      string scheme = 3;
    }

    message Notebook4 {
      string notebook_4 = 1;
    }
  }
}

message NotebookDocumentClientCapabilities {
  NotebookDocumentSyncClientCapabilities synchronization = 1;
}

message NotebookDocumentFilter {
  oneof notebook_document_filter {
    NotebookDocumentFilter1 notebook_document_filter_1 = 1;

    NotebookDocumentFilter2 notebook_document_filter_2 = 2;

    NotebookDocumentFilter3 notebook_document_filter_3 = 3;
  }

  message NotebookDocumentFilter1 {
    string scheme = 1;

    string notebook_type = 2;

    string pattern = 3;
  }

  message NotebookDocumentFilter2 {
    string notebook_type = 1;

    string pattern = 2;

    string scheme = 3;
  }

  message NotebookDocumentFilter3 {
    string notebook_type = 1;

    string pattern = 2;

    string scheme = 3;
  }
}

message NotebookDocumentSyncClientCapabilities {
  bool dynamic_registration = 1;

  bool execution_summary_support = 2;
}

message NotebookDocumentSyncOptions {
  repeated NotebookSelector notebook_selector = 1;

  bool save = 2;

  message NotebookSelector {
    NotebookSelector notebook_selector = 1;

    message NotebookSelector {
      oneof notebook_selector {
        NotebookSelector1 notebook_selector_1 = 1;

        NotebookSelector2 notebook_selector_2 = 2;
      }

      message NotebookSelector1 {
        repeated Cells cells = 1;

        Notebook notebook = 2;

        message Cells {
          Cells cells = 1;

          message Cells {
            string language = 1;
          }
        }

        message Notebook {
          oneof notebook {
            Notebook1 notebook_1 = 1;

            Notebook2 notebook_2 = 2;

            Notebook3 notebook_3 = 3;

            Notebook4 notebook_4 = 4;
          }

          message Notebook1 {
            string notebook_type = 1;

            string pattern = 2;

            string scheme = 3;
          }

          message Notebook2 {
            string notebook_type = 1;

            string pattern = 2;

            string scheme = 3;
          }

          message Notebook3 {
            string pattern = 1;

            string scheme = 2;

            string notebook_type = 3;
          }

          message Notebook4 {
            string notebook_4 = 1;
          }
        }
      }

      message NotebookSelector2 {
        repeated Cells cells = 1;

        Notebook notebook = 2;

        message Cells {
          Cells cells = 1;

          message Cells {
            string language = 1;
          }
        }

        message Notebook {
          oneof notebook {
            Notebook1 notebook_1 = 1;

            Notebook2 notebook_2 = 2;

            Notebook3 notebook_3 = 3;

            Notebook4 notebook_4 = 4;
          }

          message Notebook1 {
            string pattern = 1;

            string scheme = 2;

            string notebook_type = 3;
          }

          message Notebook2 {
            string notebook_type = 1;

            string pattern = 2;

            string scheme = 3;
          }

          message Notebook3 {
            string notebook_type = 1;

            string pattern = 2;

            string scheme = 3;
          }

          message Notebook4 {
            string notebook_4 = 1;
          }
        }
      }
    }
  }
}

message OptionalVersionedTextDocumentIdentifier {
  int32 version = 1;

  DocumentUri uri = 2;
}

message ParameterInformation {
  Label label = 1;

  Documentation documentation = 2;

  message Documentation {
    oneof documentation {
      MarkupContent markup_content = 1;

      Documentation2 documentation_2 = 2;
    }

    message Documentation2 {
      string documentation_2 = 1;
    }
  }

  message Label {
    oneof label {
      Label1 label_1 = 1;

      Label2 label_2 = 2;
    }

    message Label1 {
      Label1 label_1 = 1;

      message Label1 {
        int32 label_1 = 1;
      }
    }

    message Label2 {
      string label_2 = 1;
    }
  }
}

message PartialResultParams {
  ProgressToken partial_result_token = 1;
}

message Pattern {
  string pattern = 1;
}

message Position {
  Uinteger line = 1;

  Uinteger character = 2;
}

message PositionEncodingKind {
  string position_encoding_kind = 1;
}

message PrepareRenameParams {
  TextDocumentIdentifier text_document = 1;

  Position position = 2;

  ProgressToken work_done_token = 3;
}

message PrepareRenameResult {
  oneof prepare_rename_result {
    Range range = 1;

    PrepareRenameResult2 prepare_rename_result_2 = 2;

    PrepareRenameResult3 prepare_rename_result_3 = 3;
  }

  message PrepareRenameResult2 {
    string placeholder = 1;

    Range range = 2;
  }

  message PrepareRenameResult3 {
    bool default_behavior = 1;
  }
}

message PrepareSupportDefaultBehavior {
  enum PrepareSupportDefaultBehavior {
    PrepareSupportDefaultBehavior_1 = 1;
  }
}

message ProgressToken {
  oneof progress_token {
    ProgressToken1 progress_token_1 = 1;

    ProgressToken2 progress_token_2 = 2;
  }

  message ProgressToken1 {
    string progress_token_1 = 1;
  }

  message ProgressToken2 {
    int32 progress_token_2 = 1;
  }
}

message PublishDiagnosticsClientCapabilities {
  bool related_information = 1;

  TagSupport tag_support = 2;

  bool version_support = 3;

  bool code_description_support = 4;

  bool data_support = 5;

  message TagSupport {
    repeated ValueSet value_set = 1;

    message ValueSet {
      ValueSet value_set = 1;

      message ValueSet {
        enum ValueSet {
          ValueSet_1 = 1;

          ValueSet_2 = 2;
        }
      }
    }
  }
}

message PublishDiagnosticsParams {
  DocumentUri uri = 1;

  Integer version = 2;

  repeated Diagnostics diagnostics = 3;

  message Diagnostics {
    Diagnostic diagnostic = 1;
  }
}

message Range {
  Position start = 1;

  Position end = 2;
}

message ReferenceClientCapabilities {
  bool dynamic_registration = 1;
}

message ReferenceContext {
  bool include_declaration = 1;
}

message ReferenceOptions {
  bool work_done_progress = 1;
}

message ReferenceParams {
  ReferenceContext context = 1;

  TextDocumentIdentifier text_document = 2;

  Position position = 3;

  ProgressToken work_done_token = 4;

  ProgressToken partial_result_token = 5;
}

message ReferenceRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message Registration {
  string id = 1;

  string method = 2;
}

message RegistrationParams {
  repeated Registrations registrations = 1;

  message Registrations {
    Registration registration = 1;
  }
}

message RegularExpressionsClientCapabilities {
  string engine = 1;

  string version = 2;
}

message RelativePattern {
  BaseUri base_uri = 1;

  Pattern pattern = 2;

  message BaseUri {
    oneof base_uri {
      WorkspaceFolder workspace_folder = 1;

      BaseUri2 base_uri_2 = 2;
    }

    message BaseUri2 {
      string base_uri_2 = 1;
    }
  }
}

message RenameClientCapabilities {
  bool dynamic_registration = 1;

  bool prepare_support = 2;

  PrepareSupportDefaultBehavior prepare_support_default_behavior = 3;

  bool honors_change_annotations = 4;
}

message RenameFile {
  Kind kind = 1;

  DocumentUri old_uri = 2;

  DocumentUri new_uri = 3;

  RenameFileOptions options = 4;

  ChangeAnnotationIdentifier annotation_id = 5;

  message Kind {
    enum Kind {
      Kind_Rename = 1;
    }
  }
}

message RenameFileOptions {
  bool overwrite = 1;

  bool ignore_if_exists = 2;
}

message RenameOptions {
  bool prepare_provider = 1;

  bool work_done_progress = 2;
}

message RenameParams {
  TextDocumentIdentifier text_document = 1;

  Position position = 2;

  string new_name = 3;

  ProgressToken work_done_token = 4;
}

message RenameRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool prepare_provider = 2;

  bool work_done_progress = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message ResourceOperationKind {
  enum ResourceOperationKind {
    ResourceOperationKind_Create = 1;

    ResourceOperationKind_Delete = 2;

    ResourceOperationKind_Rename = 3;
  }
}

message SaveOptions {
  bool include_text = 1;
}

message SelectionRangeClientCapabilities {
  bool dynamic_registration = 1;
}

message SelectionRangeOptions {
  bool work_done_progress = 1;
}

message SelectionRangeRegistrationOptions {
  bool work_done_progress = 1;

  repeated DocumentSelector document_selector = 2;

  string id = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message SemanticTokensClientCapabilities {
  bool dynamic_registration = 1;

  Requests requests = 2;

  repeated TokenTypes token_types = 3;

  repeated TokenModifiers token_modifiers = 4;

  repeated Formats formats = 5;

  bool overlapping_token_support = 6;

  bool multiline_token_support = 7;

  bool server_cancel_support = 8;

  bool augments_syntax_tokens = 9;

  message Formats {
    Formats formats = 1;

    message Formats {
      enum Formats {
        Formats_Relative = 1;
      }
    }
  }

  message TokenModifiers {
    TokenModifiers token_modifiers = 1;

    message TokenModifiers {
      string token_modifiers = 1;
    }
  }

  message Requests {
    Full full = 1;

    Range range = 2;

    message Full {
      oneof full {
        Full1 full_1 = 1;

        Full2 full_2 = 2;
      }

      message Full1 {
        bool delta = 1;
      }

      message Full2 {
        bool full_2 = 1;
      }
    }

    message Range {
      oneof range {
        Range2 range_2 = 1;
      }

      message Range2 {
        bool range_2 = 1;
      }
    }
  }

  message TokenTypes {
    TokenTypes token_types = 1;

    message TokenTypes {
      string token_types = 1;
    }
  }
}

message SemanticTokensLegend {
  repeated TokenTypes token_types = 1;

  repeated TokenModifiers token_modifiers = 2;

  message TokenModifiers {
    TokenModifiers token_modifiers = 1;

    message TokenModifiers {
      string token_modifiers = 1;
    }
  }

  message TokenTypes {
    TokenTypes token_types = 1;

    message TokenTypes {
      string token_types = 1;
    }
  }
}

message SemanticTokensOptions {
  SemanticTokensLegend legend = 1;

  Range range = 2;

  Full full = 3;

  bool work_done_progress = 4;

  message Full {
    oneof full {
      Full1 full_1 = 1;

      Full2 full_2 = 2;
    }

    message Full1 {
      bool delta = 1;
    }

    message Full2 {
      bool full_2 = 1;
    }
  }

  message Range {
    oneof range {
      Range2 range_2 = 1;
    }

    message Range2 {
      bool range_2 = 1;
    }
  }
}

message SemanticTokensRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  SemanticTokensLegend legend = 2;

  Range range = 3;

  Full full = 4;

  bool work_done_progress = 5;

  string id = 6;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }

  message Full {
    oneof full {
      Full1 full_1 = 1;

      Full2 full_2 = 2;
    }

    message Full1 {
      bool delta = 1;
    }

    message Full2 {
      bool full_2 = 1;
    }
  }

  message Range {
    oneof range {
      Range2 range_2 = 1;
    }

    message Range2 {
      bool range_2 = 1;
    }
  }
}

message SemanticTokensWorkspaceClientCapabilities {
  bool refresh_support = 1;
}

message ServerCapabilities {
  PositionEncodingKind position_encoding = 1;

  TextDocumentSync text_document_sync = 2;

  NotebookDocumentSync notebook_document_sync = 3;

  CompletionOptions completion_provider = 4;

  HoverProvider hover_provider = 5;

  SignatureHelpOptions signature_help_provider = 6;

  DeclarationProvider declaration_provider = 7;

  DefinitionProvider definition_provider = 8;

  TypeDefinitionProvider type_definition_provider = 9;

  ImplementationProvider implementation_provider = 10;

  ReferencesProvider references_provider = 11;

  DocumentHighlightProvider document_highlight_provider = 12;

  DocumentSymbolProvider document_symbol_provider = 13;

  CodeActionProvider code_action_provider = 14;

  CodeLensOptions code_lens_provider = 15;

  DocumentLinkOptions document_link_provider = 16;

  ColorProvider color_provider = 17;

  WorkspaceSymbolProvider workspace_symbol_provider = 18;

  DocumentFormattingProvider document_formatting_provider = 19;

  DocumentRangeFormattingProvider document_range_formatting_provider = 20;

  DocumentOnTypeFormattingOptions document_on_type_formatting_provider = 21;

  RenameProvider rename_provider = 22;

  FoldingRangeProvider folding_range_provider = 23;

  SelectionRangeProvider selection_range_provider = 24;

  ExecuteCommandOptions execute_command_provider = 25;

  CallHierarchyProvider call_hierarchy_provider = 26;

  LinkedEditingRangeProvider linked_editing_range_provider = 27;

  SemanticTokensProvider semantic_tokens_provider = 28;

  MonikerProvider moniker_provider = 29;

  TypeHierarchyProvider type_hierarchy_provider = 30;

  InlineValueProvider inline_value_provider = 31;

  InlayHintProvider inlay_hint_provider = 32;

  Workspace workspace = 33;

  message DocumentSymbolProvider {
    oneof document_symbol_provider {
      DocumentSymbolOptions document_symbol_options = 1;

      DocumentSymbolProvider2 document_symbol_provider_2 = 2;
    }

    message DocumentSymbolProvider2 {
      bool document_symbol_provider_2 = 1;
    }
  }

  message InlayHintProvider {
    oneof inlay_hint_provider {
      InlayHintProvider3 inlay_hint_provider_3 = 1;
    }

    message InlayHintProvider3 {
      bool inlay_hint_provider_3 = 1;
    }
  }

  message SemanticTokensProvider {
    oneof semantic_tokens_provider {
      SemanticTokensOptions semantic_tokens_options = 1;

      SemanticTokensRegistrationOptions semantic_tokens_registration_options = 2;
    }
  }

  message ColorProvider {
    oneof color_provider {
      DocumentColorOptions document_color_options = 1;

      DocumentColorRegistrationOptions document_color_registration_options = 2;

      ColorProvider3 color_provider_3 = 3;
    }

    message ColorProvider3 {
      bool color_provider_3 = 1;
    }
  }

  message ImplementationProvider {
    oneof implementation_provider {
      ImplementationOptions implementation_options = 1;

      ImplementationRegistrationOptions implementation_registration_options = 2;

      ImplementationProvider3 implementation_provider_3 = 3;
    }

    message ImplementationProvider3 {
      bool implementation_provider_3 = 1;
    }
  }

  message MonikerProvider {
    oneof moniker_provider {
      MonikerOptions moniker_options = 1;

      MonikerRegistrationOptions moniker_registration_options = 2;

      MonikerProvider3 moniker_provider_3 = 3;
    }

    message MonikerProvider3 {
      bool moniker_provider_3 = 1;
    }
  }

  message TextDocumentSync {
    oneof text_document_sync {
      TextDocumentSyncOptions text_document_sync_options = 1;

      TextDocumentSync2 text_document_sync_2 = 2;
    }

    message TextDocumentSync2 {
      enum TextDocumentSync2 {
        TextDocumentSync2_0 = 1;

        TextDocumentSync2_1 = 2;

        TextDocumentSync2_2 = 3;
      }
    }
  }

  message WorkspaceSymbolProvider {
    oneof workspace_symbol_provider {
      WorkspaceSymbolOptions workspace_symbol_options = 1;

      WorkspaceSymbolProvider2 workspace_symbol_provider_2 = 2;
    }

    message WorkspaceSymbolProvider2 {
      bool workspace_symbol_provider_2 = 1;
    }
  }

  message FoldingRangeProvider {
    oneof folding_range_provider {
      FoldingRangeOptions folding_range_options = 1;

      FoldingRangeRegistrationOptions folding_range_registration_options = 2;

      FoldingRangeProvider3 folding_range_provider_3 = 3;
    }

    message FoldingRangeProvider3 {
      bool folding_range_provider_3 = 1;
    }
  }

  message NotebookDocumentSync {
    oneof notebook_document_sync {
      NotebookDocumentSyncOptions notebook_document_sync_options = 1;
    }
  }

  message DocumentHighlightProvider {
    oneof document_highlight_provider {
      DocumentHighlightOptions document_highlight_options = 1;

      DocumentHighlightProvider2 document_highlight_provider_2 = 2;
    }

    message DocumentHighlightProvider2 {
      bool document_highlight_provider_2 = 1;
    }
  }

  message HoverProvider {
    oneof hover_provider {
      HoverOptions hover_options = 1;

      HoverProvider2 hover_provider_2 = 2;
    }

    message HoverProvider2 {
      bool hover_provider_2 = 1;
    }
  }

  message InlineValueProvider {
    oneof inline_value_provider {
      WorkDoneProgressOptions work_done_progress_options = 1;

      InlineValueProvider3 inline_value_provider_3 = 2;
    }

    message InlineValueProvider3 {
      bool inline_value_provider_3 = 1;
    }
  }

  message TypeDefinitionProvider {
    oneof type_definition_provider {
      TypeDefinitionOptions type_definition_options = 1;

      TypeDefinitionRegistrationOptions type_definition_registration_options = 2;

      TypeDefinitionProvider3 type_definition_provider_3 = 3;
    }

    message TypeDefinitionProvider3 {
      bool type_definition_provider_3 = 1;
    }
  }

  message TypeHierarchyProvider {
    oneof type_hierarchy_provider {
      WorkDoneProgressOptions work_done_progress_options = 1;

      TypeHierarchyProvider3 type_hierarchy_provider_3 = 2;
    }

    message TypeHierarchyProvider3 {
      bool type_hierarchy_provider_3 = 1;
    }
  }

  message DeclarationProvider {
    oneof declaration_provider {
      DeclarationOptions declaration_options = 1;

      DeclarationRegistrationOptions declaration_registration_options = 2;

      DeclarationProvider3 declaration_provider_3 = 3;
    }

    message DeclarationProvider3 {
      bool declaration_provider_3 = 1;
    }
  }

  message DocumentRangeFormattingProvider {
    oneof document_range_formatting_provider {
      DocumentRangeFormattingOptions document_range_formatting_options = 1;

      DocumentRangeFormattingProvider2 document_range_formatting_provider_2 = 2;
    }

    message DocumentRangeFormattingProvider2 {
      bool document_range_formatting_provider_2 = 1;
    }
  }

  message ReferencesProvider {
    oneof references_provider {
      ReferenceOptions reference_options = 1;

      ReferencesProvider2 references_provider_2 = 2;
    }

    message ReferencesProvider2 {
      bool references_provider_2 = 1;
    }
  }

  message SelectionRangeProvider {
    oneof selection_range_provider {
      SelectionRangeOptions selection_range_options = 1;

      SelectionRangeRegistrationOptions selection_range_registration_options = 2;

      SelectionRangeProvider3 selection_range_provider_3 = 3;
    }

    message SelectionRangeProvider3 {
      bool selection_range_provider_3 = 1;
    }
  }

  message CodeActionProvider {
    oneof code_action_provider {
      CodeActionOptions code_action_options = 1;

      CodeActionProvider2 code_action_provider_2 = 2;
    }

    message CodeActionProvider2 {
      bool code_action_provider_2 = 1;
    }
  }

  message DefinitionProvider {
    oneof definition_provider {
      DefinitionOptions definition_options = 1;

      DefinitionProvider2 definition_provider_2 = 2;
    }

    message DefinitionProvider2 {
      bool definition_provider_2 = 1;
    }
  }

  message DocumentFormattingProvider {
    oneof document_formatting_provider {
      DocumentFormattingOptions document_formatting_options = 1;

      DocumentFormattingProvider2 document_formatting_provider_2 = 2;
    }

    message DocumentFormattingProvider2 {
      bool document_formatting_provider_2 = 1;
    }
  }

  message LinkedEditingRangeProvider {
    oneof linked_editing_range_provider {
      LinkedEditingRangeOptions linked_editing_range_options = 1;

      LinkedEditingRangeRegistrationOptions linked_editing_range_registration_options = 2;

      LinkedEditingRangeProvider3 linked_editing_range_provider_3 = 3;
    }

    message LinkedEditingRangeProvider3 {
      bool linked_editing_range_provider_3 = 1;
    }
  }

  message RenameProvider {
    oneof rename_provider {
      RenameOptions rename_options = 1;

      RenameProvider2 rename_provider_2 = 2;
    }

    message RenameProvider2 {
      bool rename_provider_2 = 1;
    }
  }

  message Workspace {
    WorkspaceFoldersServerCapabilities workspace_folders = 1;

    FileOperationOptions file_operations = 2;
  }

  message CallHierarchyProvider {
    oneof call_hierarchy_provider {
      CallHierarchyOptions call_hierarchy_options = 1;

      CallHierarchyRegistrationOptions call_hierarchy_registration_options = 2;

      CallHierarchyProvider3 call_hierarchy_provider_3 = 3;
    }

    message CallHierarchyProvider3 {
      bool call_hierarchy_provider_3 = 1;
    }
  }
}

message ServerCapabilitiesT {
  PositionEncodingKind position_encoding = 1;

  TextDocumentSync text_document_sync = 2;

  NotebookDocumentSync notebook_document_sync = 3;

  CompletionOptions completion_provider = 4;

  HoverProvider hover_provider = 5;

  SignatureHelpOptions signature_help_provider = 6;

  DeclarationProvider declaration_provider = 7;

  DefinitionProvider definition_provider = 8;

  TypeDefinitionProvider type_definition_provider = 9;

  ImplementationProvider implementation_provider = 10;

  ReferencesProvider references_provider = 11;

  DocumentHighlightProvider document_highlight_provider = 12;

  DocumentSymbolProvider document_symbol_provider = 13;

  CodeActionProvider code_action_provider = 14;

  CodeLensOptions code_lens_provider = 15;

  DocumentLinkOptions document_link_provider = 16;

  ColorProvider color_provider = 17;

  WorkspaceSymbolProvider workspace_symbol_provider = 18;

  DocumentFormattingProvider document_formatting_provider = 19;

  DocumentRangeFormattingProvider document_range_formatting_provider = 20;

  DocumentOnTypeFormattingOptions document_on_type_formatting_provider = 21;

  RenameProvider rename_provider = 22;

  FoldingRangeProvider folding_range_provider = 23;

  SelectionRangeProvider selection_range_provider = 24;

  ExecuteCommandOptions execute_command_provider = 25;

  CallHierarchyProvider call_hierarchy_provider = 26;

  LinkedEditingRangeProvider linked_editing_range_provider = 27;

  SemanticTokensProvider semantic_tokens_provider = 28;

  MonikerProvider moniker_provider = 29;

  TypeHierarchyProvider type_hierarchy_provider = 30;

  InlineValueProvider inline_value_provider = 31;

  InlayHintProvider inlay_hint_provider = 32;

  Workspace workspace = 33;

  message DocumentSymbolProvider {
    oneof document_symbol_provider {
      DocumentSymbolOptions document_symbol_options = 1;

      DocumentSymbolProvider2 document_symbol_provider_2 = 2;
    }

    message DocumentSymbolProvider2 {
      bool document_symbol_provider_2 = 1;
    }
  }

  message LinkedEditingRangeProvider {
    oneof linked_editing_range_provider {
      LinkedEditingRangeOptions linked_editing_range_options = 1;

      LinkedEditingRangeRegistrationOptions linked_editing_range_registration_options = 2;

      LinkedEditingRangeProvider3 linked_editing_range_provider_3 = 3;
    }

    message LinkedEditingRangeProvider3 {
      bool linked_editing_range_provider_3 = 1;
    }
  }

  message NotebookDocumentSync {
    oneof notebook_document_sync {
      NotebookDocumentSyncOptions notebook_document_sync_options = 1;
    }
  }

  message DefinitionProvider {
    oneof definition_provider {
      DefinitionOptions definition_options = 1;

      DefinitionProvider2 definition_provider_2 = 2;
    }

    message DefinitionProvider2 {
      bool definition_provider_2 = 1;
    }
  }

  message TextDocumentSync {
    oneof text_document_sync {
      TextDocumentSyncOptions text_document_sync_options = 1;

      TextDocumentSync2 text_document_sync_2 = 2;
    }

    message TextDocumentSync2 {
      enum TextDocumentSync2 {
        TextDocumentSync2_0 = 1;

        TextDocumentSync2_1 = 2;

        TextDocumentSync2_2 = 3;
      }
    }
  }

  message WorkspaceSymbolProvider {
    oneof workspace_symbol_provider {
      WorkspaceSymbolOptions workspace_symbol_options = 1;

      WorkspaceSymbolProvider2 workspace_symbol_provider_2 = 2;
    }

    message WorkspaceSymbolProvider2 {
      bool workspace_symbol_provider_2 = 1;
    }
  }

  message DocumentFormattingProvider {
    oneof document_formatting_provider {
      DocumentFormattingOptions document_formatting_options = 1;

      DocumentFormattingProvider2 document_formatting_provider_2 = 2;
    }

    message DocumentFormattingProvider2 {
      bool document_formatting_provider_2 = 1;
    }
  }

  message DocumentHighlightProvider {
    oneof document_highlight_provider {
      DocumentHighlightOptions document_highlight_options = 1;

      DocumentHighlightProvider2 document_highlight_provider_2 = 2;
    }

    message DocumentHighlightProvider2 {
      bool document_highlight_provider_2 = 1;
    }
  }

  message FoldingRangeProvider {
    oneof folding_range_provider {
      FoldingRangeOptions folding_range_options = 1;

      FoldingRangeRegistrationOptions folding_range_registration_options = 2;

      FoldingRangeProvider3 folding_range_provider_3 = 3;
    }

    message FoldingRangeProvider3 {
      bool folding_range_provider_3 = 1;
    }
  }

  message HoverProvider {
    oneof hover_provider {
      HoverOptions hover_options = 1;

      HoverProvider2 hover_provider_2 = 2;
    }

    message HoverProvider2 {
      bool hover_provider_2 = 1;
    }
  }

  message ReferencesProvider {
    oneof references_provider {
      ReferenceOptions reference_options = 1;

      ReferencesProvider2 references_provider_2 = 2;
    }

    message ReferencesProvider2 {
      bool references_provider_2 = 1;
    }
  }

  message TypeHierarchyProvider {
    oneof type_hierarchy_provider {
      WorkDoneProgressOptions work_done_progress_options = 1;

      TypeHierarchyProvider3 type_hierarchy_provider_3 = 2;
    }

    message TypeHierarchyProvider3 {
      bool type_hierarchy_provider_3 = 1;
    }
  }

  message ColorProvider {
    oneof color_provider {
      DocumentColorOptions document_color_options = 1;

      DocumentColorRegistrationOptions document_color_registration_options = 2;

      ColorProvider3 color_provider_3 = 3;
    }

    message ColorProvider3 {
      bool color_provider_3 = 1;
    }
  }

  message DeclarationProvider {
    oneof declaration_provider {
      DeclarationOptions declaration_options = 1;

      DeclarationRegistrationOptions declaration_registration_options = 2;

      DeclarationProvider3 declaration_provider_3 = 3;
    }

    message DeclarationProvider3 {
      bool declaration_provider_3 = 1;
    }
  }

  message TypeDefinitionProvider {
    oneof type_definition_provider {
      TypeDefinitionOptions type_definition_options = 1;

      TypeDefinitionRegistrationOptions type_definition_registration_options = 2;

      TypeDefinitionProvider3 type_definition_provider_3 = 3;
    }

    message TypeDefinitionProvider3 {
      bool type_definition_provider_3 = 1;
    }
  }

  message CallHierarchyProvider {
    oneof call_hierarchy_provider {
      CallHierarchyOptions call_hierarchy_options = 1;

      CallHierarchyRegistrationOptions call_hierarchy_registration_options = 2;

      CallHierarchyProvider3 call_hierarchy_provider_3 = 3;
    }

    message CallHierarchyProvider3 {
      bool call_hierarchy_provider_3 = 1;
    }
  }

  message CodeActionProvider {
    oneof code_action_provider {
      CodeActionOptions code_action_options = 1;

      CodeActionProvider2 code_action_provider_2 = 2;
    }

    message CodeActionProvider2 {
      bool code_action_provider_2 = 1;
    }
  }

  message RenameProvider {
    oneof rename_provider {
      RenameOptions rename_options = 1;

      RenameProvider2 rename_provider_2 = 2;
    }

    message RenameProvider2 {
      bool rename_provider_2 = 1;
    }
  }

  message Workspace {
    FileOperationOptions file_operations = 1;

    WorkspaceFoldersServerCapabilities workspace_folders = 2;
  }

  message InlineValueProvider {
    oneof inline_value_provider {
      WorkDoneProgressOptions work_done_progress_options = 1;

      InlineValueProvider3 inline_value_provider_3 = 2;
    }

    message InlineValueProvider3 {
      bool inline_value_provider_3 = 1;
    }
  }

  message InlayHintProvider {
    oneof inlay_hint_provider {
      InlayHintProvider3 inlay_hint_provider_3 = 1;
    }

    message InlayHintProvider3 {
      bool inlay_hint_provider_3 = 1;
    }
  }

  message MonikerProvider {
    oneof moniker_provider {
      MonikerOptions moniker_options = 1;

      MonikerRegistrationOptions moniker_registration_options = 2;

      MonikerProvider3 moniker_provider_3 = 3;
    }

    message MonikerProvider3 {
      bool moniker_provider_3 = 1;
    }
  }

  message SelectionRangeProvider {
    oneof selection_range_provider {
      SelectionRangeOptions selection_range_options = 1;

      SelectionRangeRegistrationOptions selection_range_registration_options = 2;

      SelectionRangeProvider3 selection_range_provider_3 = 3;
    }

    message SelectionRangeProvider3 {
      bool selection_range_provider_3 = 1;
    }
  }

  message DocumentRangeFormattingProvider {
    oneof document_range_formatting_provider {
      DocumentRangeFormattingOptions document_range_formatting_options = 1;

      DocumentRangeFormattingProvider2 document_range_formatting_provider_2 = 2;
    }

    message DocumentRangeFormattingProvider2 {
      bool document_range_formatting_provider_2 = 1;
    }
  }

  message ImplementationProvider {
    oneof implementation_provider {
      ImplementationOptions implementation_options = 1;

      ImplementationRegistrationOptions implementation_registration_options = 2;

      ImplementationProvider3 implementation_provider_3 = 3;
    }

    message ImplementationProvider3 {
      bool implementation_provider_3 = 1;
    }
  }

  message SemanticTokensProvider {
    oneof semantic_tokens_provider {
      SemanticTokensOptions semantic_tokens_options = 1;

      SemanticTokensRegistrationOptions semantic_tokens_registration_options = 2;
    }
  }
}

message ShowDocumentClientCapabilities {
  bool support = 1;
}

message ShowMessageParams {
  MessageType type = 1;

  string message = 2;
}

message ShowMessageRequestClientCapabilities {
  bool message_action_item = 1;
}

message ShowMessageRequestParams {
  MessageType type = 1;

  string message = 2;

  repeated Actions actions = 3;

  message Actions {
    MessageActionItem message_action_item = 1;
  }
}

message SignatureHelp {
  repeated Signatures signatures = 1;

  Uinteger active_signature = 2;

  Uinteger active_parameter = 3;

  message Signatures {
    SignatureInformation signature_information = 1;
  }
}

message SignatureHelpClientCapabilities {
  bool dynamic_registration = 1;

  SignatureInformation signature_information = 2;

  bool context_support = 3;

  message SignatureInformation {
    bool active_parameter_support = 1;

    repeated DocumentationFormat documentation_format = 2;

    bool parameter_information = 3;

    message DocumentationFormat {
      DocumentationFormat documentation_format = 1;

      message DocumentationFormat {
        enum DocumentationFormat {
          DocumentationFormat_Markdown = 1;

          DocumentationFormat_Plaintext = 2;
        }
      }
    }
  }
}

message SignatureHelpContext {
  SignatureHelpTriggerKind trigger_kind = 1;

  string trigger_character = 2;

  bool is_retrigger = 3;

  SignatureHelp active_signature_help = 4;
}

message SignatureHelpOptions {
  repeated TriggerCharacters trigger_characters = 1;

  repeated RetriggerCharacters retrigger_characters = 2;

  bool work_done_progress = 3;

  message RetriggerCharacters {
    RetriggerCharacters retrigger_characters = 1;

    message RetriggerCharacters {
      string retrigger_characters = 1;
    }
  }

  message TriggerCharacters {
    TriggerCharacters trigger_characters = 1;

    message TriggerCharacters {
      string trigger_characters = 1;
    }
  }
}

message SignatureHelpParams {
  SignatureHelpContext context = 1;

  TextDocumentIdentifier text_document = 2;

  Position position = 3;

  ProgressToken work_done_token = 4;
}

message SignatureHelpRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  repeated TriggerCharacters trigger_characters = 2;

  repeated RetriggerCharacters retrigger_characters = 3;

  bool work_done_progress = 4;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }

  message RetriggerCharacters {
    RetriggerCharacters retrigger_characters = 1;

    message RetriggerCharacters {
      string retrigger_characters = 1;
    }
  }

  message TriggerCharacters {
    TriggerCharacters trigger_characters = 1;

    message TriggerCharacters {
      string trigger_characters = 1;
    }
  }
}

message SignatureHelpTriggerKind {
  enum SignatureHelpTriggerKind {
    SignatureHelpTriggerKind_1 = 1;

    SignatureHelpTriggerKind_2 = 2;

    SignatureHelpTriggerKind_3 = 3;
  }
}

message SignatureInformation {
  string label = 1;

  Documentation documentation = 2;

  repeated Parameters parameters = 3;

  Uinteger active_parameter = 4;

  message Documentation {
    oneof documentation {
      MarkupContent markup_content = 1;

      Documentation2 documentation_2 = 2;
    }

    message Documentation2 {
      string documentation_2 = 1;
    }
  }

  message Parameters {
    ParameterInformation parameter_information = 1;
  }
}

message StaticRegistrationOptions {
  string id = 1;
}

message TextDocumentChangeRegistrationOptions {
  TextDocumentSyncKind sync_kind = 1;

  repeated DocumentSelector document_selector = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message TextDocumentClientCapabilities {
  TextDocumentSyncClientCapabilities synchronization = 1;

  CompletionClientCapabilities completion = 2;

  HoverClientCapabilities hover = 3;

  SignatureHelpClientCapabilities signature_help = 4;

  DeclarationClientCapabilities declaration = 5;

  DefinitionClientCapabilities definition = 6;

  TypeDefinitionClientCapabilities type_definition = 7;

  ImplementationClientCapabilities implementation = 8;

  ReferenceClientCapabilities references = 9;

  DocumentHighlightClientCapabilities document_highlight = 10;

  DocumentSymbolClientCapabilities document_symbol = 11;

  CodeActionClientCapabilities code_action = 12;

  CodeLensClientCapabilities code_lens = 13;

  DocumentLinkClientCapabilities document_link = 14;

  DocumentColorClientCapabilities color_provider = 15;

  DocumentFormattingClientCapabilities formatting = 16;

  DocumentRangeFormattingClientCapabilities range_formatting = 17;

  DocumentOnTypeFormattingClientCapabilities on_type_formatting = 18;

  RenameClientCapabilities rename = 19;

  FoldingRangeClientCapabilities folding_range = 20;

  SelectionRangeClientCapabilities selection_range = 21;

  PublishDiagnosticsClientCapabilities publish_diagnostics = 22;

  CallHierarchyClientCapabilities call_hierarchy = 23;

  SemanticTokensClientCapabilities semantic_tokens = 24;

  LinkedEditingRangeClientCapabilities linked_editing_range = 25;

  MonikerClientCapabilities moniker = 26;

  TypeHierarchyClientCapabilities type_hierarchy = 27;

  InlineValueClientCapabilities inline_value = 28;

  InlayHintClientCapabilities inlay_hint = 29;

  DiagnosticClientCapabilities diagnostic = 30;
}

message TextDocumentContentChangeEvent {
  oneof text_document_content_change_event {
    TextDocumentContentChangeEvent1 text_document_content_change_event_1 = 1;

    TextDocumentContentChangeEvent2 text_document_content_change_event_2 = 2;
  }

  message TextDocumentContentChangeEvent1 {
    Range range = 1;

    Uinteger range_length = 2;

    string text = 3;
  }

  message TextDocumentContentChangeEvent2 {
    string text = 1;
  }
}

message TextDocumentEdit {
  OptionalVersionedTextDocumentIdentifier text_document = 1;

  repeated Edits edits = 2;

  message Edits {
    Edits edits = 1;

    message Edits {
      oneof edits {
        TextEdit text_edit = 1;

        AnnotatedTextEdit annotated_text_edit = 2;
      }
    }
  }
}

message TextDocumentFilter {
  oneof text_document_filter {
    TextDocumentFilter1 text_document_filter_1 = 1;

    TextDocumentFilter2 text_document_filter_2 = 2;

    TextDocumentFilter3 text_document_filter_3 = 3;
  }

  message TextDocumentFilter1 {
    string language = 1;

    string pattern = 2;

    string scheme = 3;
  }

  message TextDocumentFilter2 {
    string scheme = 1;

    string language = 2;

    string pattern = 3;
  }

  message TextDocumentFilter3 {
    string scheme = 1;

    string language = 2;

    string pattern = 3;
  }
}

message TextDocumentIdentifier {
  DocumentUri uri = 1;
}

message TextDocumentItem {
  DocumentUri uri = 1;

  string language_id = 2;

  Integer version = 3;

  string text = 4;
}

message TextDocumentPositionParams {
  TextDocumentIdentifier text_document = 1;

  Position position = 2;
}

message TextDocumentRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message TextDocumentSaveReason {
  enum TextDocumentSaveReason {
    TextDocumentSaveReason_1 = 1;

    TextDocumentSaveReason_2 = 2;

    TextDocumentSaveReason_3 = 3;
  }
}

message TextDocumentSaveRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool include_text = 2;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message TextDocumentSyncClientCapabilities {
  bool dynamic_registration = 1;

  bool will_save = 2;

  bool will_save_wait_until = 3;

  bool did_save = 4;
}

message TextDocumentSyncKind {
  enum TextDocumentSyncKind {
    TextDocumentSyncKind_0 = 1;

    TextDocumentSyncKind_1 = 2;

    TextDocumentSyncKind_2 = 3;
  }
}

message TextDocumentSyncOptions {
  bool open_close = 1;

  TextDocumentSyncKind change = 2;

  bool will_save = 3;

  bool will_save_wait_until = 4;

  Save save = 5;

  message Save {
    oneof save {
      SaveOptions save_options = 1;

      Save2 save_2 = 2;
    }

    message Save2 {
      bool save_2 = 1;
    }
  }
}

message TextEdit {
  Range range = 1;

  string new_text = 2;
}

message TypeDefinitionClientCapabilities {
  bool dynamic_registration = 1;

  bool link_support = 2;
}

message TypeDefinitionOptions {
  bool work_done_progress = 1;
}

message TypeDefinitionRegistrationOptions {
  repeated DocumentSelector document_selector = 1;

  bool work_done_progress = 2;

  string id = 3;

  message DocumentSelector {
    DocumentSelector document_selector = 1;
  }
}

message TypeHierarchyClientCapabilities {
  bool dynamic_registration = 1;
}

message URI {
  string uri = 1;
}

message Uinteger {
  int32 uinteger = 1;
}

message Unregistration {
  string id = 1;

  string method = 2;
}

message UnregistrationParams {
  repeated Unregisterations unregisterations = 1;

  message Unregisterations {
    Unregistration unregistration = 1;
  }
}

message VersionedTextDocumentIdentifier {
  Integer version = 1;

  DocumentUri uri = 2;
}

message WatchKind {
  int32 watch_kind = 1;
}

message WillSaveTextDocumentParams {
  TextDocumentIdentifier text_document = 1;

  TextDocumentSaveReason reason = 2;
}

message WindowClientCapabilities {
  bool work_done_progress = 1;

  ShowMessageRequestClientCapabilities show_message = 2;

  ShowDocumentClientCapabilities show_document = 3;
}

message WorkDoneProgressOptions {
  bool work_done_progress = 1;
}

message WorkDoneProgressParams {
  ProgressToken work_done_token = 1;
}

message WorkspaceClientCapabilities {
  bool apply_edit = 1;

  WorkspaceEditClientCapabilities workspace_edit = 2;

  DidChangeConfigurationClientCapabilities did_change_configuration = 3;

  DidChangeWatchedFilesClientCapabilities did_change_watched_files = 4;

  WorkspaceSymbolClientCapabilities symbol = 5;

  ExecuteCommandClientCapabilities execute_command = 6;

  bool workspace_folders = 7;

  bool configuration = 8;

  SemanticTokensWorkspaceClientCapabilities semantic_tokens = 9;

  CodeLensWorkspaceClientCapabilities code_lens = 10;

  FileOperationClientCapabilities file_operations = 11;

  InlineValueWorkspaceClientCapabilities inline_value = 12;

  InlayHintWorkspaceClientCapabilities inlay_hint = 13;

  DiagnosticWorkspaceClientCapabilities diagnostics = 14;
}

message WorkspaceEdit {
  AdditionalProperties changes = 1;

  repeated DocumentChanges document_changes = 2;

  ChangeAnnotation change_annotations = 3;

  message ChangeAnnotation {
    bool needs_confirmation = 1;

    string description = 2;

    string label = 3;
  }

  message AdditionalProperties {
    TextEdit text_edit = 1;
  }

  message DocumentChanges {
    DocumentChanges document_changes = 1;

    message DocumentChanges {
      oneof document_changes {
        TextDocumentEdit text_document_edit = 1;

        CreateFile create_file = 2;

        RenameFile rename_file = 3;

        DeleteFile delete_file = 4;
      }
    }
  }
}

message WorkspaceEditClientCapabilities {
  bool document_changes = 1;

  repeated ResourceOperations resource_operations = 2;

  FailureHandlingKind failure_handling = 3;

  bool normalizes_line_endings = 4;

  bool change_annotation_support = 5;

  message ResourceOperations {
    ResourceOperations resource_operations = 1;

    message ResourceOperations {
      enum ResourceOperations {
        ResourceOperations_Create = 1;

        ResourceOperations_Delete = 2;

        ResourceOperations_Rename = 3;
      }
    }
  }
}

message WorkspaceFolder {
  URI uri = 1;

  string name = 2;
}

message WorkspaceFoldersInitializeParams {
  repeated WorkspaceFolders workspace_folders = 1;

  message WorkspaceFolders {
    WorkspaceFolder workspace_folder = 1;
  }
}

message WorkspaceFoldersServerCapabilities {
  bool supported = 1;

  ChangeNotifications change_notifications = 2;

  message ChangeNotifications {
    oneof change_notifications {
      ChangeNotifications1 change_notifications_1 = 1;

      ChangeNotifications2 change_notifications_2 = 2;
    }

    message ChangeNotifications1 {
      string change_notifications_1 = 1;
    }

    message ChangeNotifications2 {
      bool change_notifications_2 = 1;
    }
  }
}

message WorkspaceSymbolClientCapabilities {
  bool dynamic_registration = 1;

  SymbolKind symbol_kind = 2;

  TagSupport tag_support = 3;

  ResolveSupport resolve_support = 4;

  message ResolveSupport {
    repeated Properties properties = 1;

    message Properties {
      Properties properties = 1;

      message Properties {
        string properties = 1;
      }
    }
  }

  message SymbolKind {
    repeated ValueSet value_set = 1;

    message ValueSet {
      ValueSet value_set = 1;

      message ValueSet {
        enum ValueSet {
          ValueSet_1 = 1;

          ValueSet_2 = 2;

          ValueSet_3 = 3;

          ValueSet_4 = 4;

          ValueSet_5 = 5;

          ValueSet_6 = 6;

          ValueSet_7 = 7;

          ValueSet_8 = 8;

          ValueSet_9 = 9;

          ValueSet_10 = 10;

          ValueSet_11 = 11;

          ValueSet_12 = 12;

          ValueSet_13 = 13;

          ValueSet_14 = 14;

          ValueSet_15 = 15;

          ValueSet_16 = 16;

          ValueSet_17 = 17;

          ValueSet_18 = 18;

          ValueSet_19 = 19;

          ValueSet_20 = 20;

          ValueSet_21 = 21;

          ValueSet_22 = 22;

          ValueSet_23 = 23;

          ValueSet_24 = 24;

          ValueSet_25 = 25;

          ValueSet_26 = 26;
        }
      }
    }
  }

  message TagSupport {
    repeated ValueSet value_set = 1;

    message ValueSet {
      ValueSet value_set = 1;

      message ValueSet {
        enum ValueSet {
          ValueSet_1 = 1;
        }
      }
    }
  }
}

message WorkspaceSymbolOptions {
  bool resolve_provider = 1;

  bool work_done_progress = 2;
}

message WorkspaceSymbolParams {
  string query = 1;

  ProgressToken work_done_token = 2;

  ProgressToken partial_result_token = 3;
}

message WorkspaceSymbolRegistrationOptions {
  bool resolve_provider = 1;

  bool work_done_progress = 2;
}
```

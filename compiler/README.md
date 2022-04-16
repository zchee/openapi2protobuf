```proto
syntax = "proto3";

import "{generated-file-0001}.proto";

import "{generated-file-0002}.proto";

import "{generated-file-0003}.proto";

import "{generated-file-0004}.proto";

import "{generated-file-0005}.proto";

import "{generated-file-0006}.proto";

import "{generated-file-0007}.proto";

import "{generated-file-0008}.proto";

import "{generated-file-0009}.proto";

message CharCode {
  CharCode char_code = 1;

  enum CharCode {
    CharCode_10 = 0;

    CharCode_13 = 1;
  }
}

// DocumentUri a tagging type for string properties that are actually URIs.
message DocumentUri {
  string document_uri = 1;
}

message FullTextDocument {
  Content _content = 1;

  LanguageId language_id = 2;

  Uri uri = 3;

  LanguageId _language_id = 4;

  repeated LineOffsets _line_offsets = 5;

  DocumentUri _uri = 6;

  Version _version = 7;

  LineCount line_count = 8;

  Version version = 9;

  message Content {
    string _content = 1;
  }

  message LanguageId {
    string language_id = 1;
  }

  message Uri {
    string uri = 1;
  }

  message LineOffsets {
    LineOffsets _line_offsets = 1;
  }

  message Version {
    int32 _version = 1;
  }

  message LineCount {
    int32 line_count = 1;
  }
}

message Position {
  Character character = 1;

  Line line = 2;

  // Character character offset on a line in a document (zero-based). Assuming that the line is represented as a string, the `character` value represents the gap between the `character` and `character + 1`.  If the character value is greater than the line length it defaults back to the line length. If a line number is negative, it defaults to 0.
  message Character {
    int32 character = 1;
  }

  // Line line position in a document (zero-based). If a line number is greater than the number of lines in a document, it defaults back to the number of lines in the document. If a line number is negative, it defaults to 0.
  message Line {
    int32 line = 1;
  }
}

message Range {
  Position end = 1;

  Position start = 2;
}

message TextDocument {
  LanguageId language_id = 1;

  LineCount line_count = 2;

  DocumentUri uri = 3;

  Version version = 4;

  // LanguageId the identifier of the language associated with this document.
  message LanguageId {
    string language_id = 1;
  }

  // LineCount the number of lines in this document.
  message LineCount {
    int32 line_count = 1;
  }

  // Version the version number of this document (it will increase after each change, including undo/redo).
  message Version {
    int32 version = 1;
  }
}

message TextDocumentContentChangeEvent {
  oneof text_document_content_change_event {
    TextDocumentContentChangeEvent_0 text_document_content_change_event_0 = 1;

    TextDocumentContentChangeEvent_1 text_document_content_change_event_1 = 2;
  }

  message TextDocumentContentChangeEvent_0 {
    RangeLength range_length = 1;

    Text text = 2;

    Range range = 3;

    // RangeLength the optional length of the range that got replaced.
    message RangeLength {
      int32 range_length = 1;
    }

    // Text the new text for the provided range.
    message Text {
      string text = 1;
    }
  }

  message TextDocumentContentChangeEvent_1 {
    Text text = 1;

    // Text the new text of the whole document.
    message Text {
      string text = 1;
    }
  }
}

message TextEdit {
  NewText new_text = 1;

  Range range = 2;

  // NewText the string to be inserted. For delete operations use an empty string.
  message NewText {
    string new_text = 1;
  }
}
```

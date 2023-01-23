module default {
  scalar type TodoNumber extending sequence;

  type Todo {
    property no -> TodoNumber;
    required property title -> str;
    property body -> str;
    required property status -> bool;
    property tag -> array<str>;
  }

  type Account {
    property no -> TodoNumber;

    property create_at -> datetime;
    property update_at -> datetime;
    property delete_at -> datetime;

    required property name -> str;
    required property currency -> str;
    required property balance -> float64 {
      default := 0.0;
    }
    required property status -> bool;
  }
}

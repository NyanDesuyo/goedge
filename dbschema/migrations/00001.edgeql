CREATE MIGRATION m1ifgd736hzpasiiscul5a3ikzwgqrkuoaq5v7sth6l6tdtvbiwgyq
    ONTO initial
{
  CREATE FUTURE nonrecursive_access_policies;
  CREATE SCALAR TYPE default::TodoNumber EXTENDING std::sequence;
  CREATE TYPE default::Account {
      CREATE REQUIRED PROPERTY balance -> std::float64 {
          SET default := 0.0;
      };
      CREATE PROPERTY create_at -> std::datetime;
      CREATE REQUIRED PROPERTY currency -> std::str;
      CREATE PROPERTY delete_at -> std::datetime;
      CREATE REQUIRED PROPERTY name -> std::str;
      CREATE PROPERTY no -> default::TodoNumber;
      CREATE REQUIRED PROPERTY status -> std::bool;
      CREATE PROPERTY update_at -> std::datetime;
  };
  CREATE TYPE default::Todo {
      CREATE PROPERTY body -> std::str;
      CREATE PROPERTY no -> default::TodoNumber;
      CREATE REQUIRED PROPERTY status -> std::bool;
      CREATE PROPERTY tag -> array<std::str>;
      CREATE REQUIRED PROPERTY title -> std::str;
  };
};

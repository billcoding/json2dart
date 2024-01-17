# json2dart
The JSON to Dart class converter

```dart
/*
Generated by json2dart (https://github.com/billcoding/json2dart)
Created at 2024-01-17T19:42:10
-------------------------------------------------
{
  "intV": 100,
  "intArrV": [100],
  "doubleV": 100.25,
  "doubleArrV": [100.25],
  "stringV": "helo",
  "stringArrV": ["helo"],
  "boolV": false,
  "boolArrV": [false],
  "objV": {
    "intV": 100,
    "intArrV": [100],
    "doubleV": 100.25,
    "doubleArrV": [100.25],
    "stringV": "helo",
    "stringArrV": ["helo"],
    "boolV": false,
    "boolArrV": [false]
  },
  "objArrV": [{
    "intV": 100,
    "intArrV": [100],
    "doubleV": 100.25,
    "doubleArrV": [100.25],
    "stringV": "helo",
    "stringArrV": ["helo"],
    "boolV": false,
    "boolArrV": [false]
  }]
}
*/

class Json2Dart {
  final List<bool> boolArrV;
  final bool boolV;
  final List<double> doubleArrV;
  final double doubleV;
  final List<int> intArrV;
  final int intV;
  final List<Json2DartObjArrV> objArrV;
  final Json2DartObjV objV;
  final List<String> stringArrV;
  final String stringV;

  const Json2Dart(this.boolArrV, this.boolV, this.doubleArrV, this.doubleV, this.intArrV, this.intV, this.objArrV, this.objV, this.stringArrV, this.stringV);

  factory Json2Dart.fromMap(Map<String, dynamic> m) {
    return Json2Dart(
      m.containsKey("boolArrV") && m["boolArrV"] != null ? (m["boolArrV"] as List<dynamic>).map((a) => a as bool).toList(): [],
      m.containsKey("boolV") && m["boolV"] != null ? m["boolV"]: false,
      m.containsKey("doubleArrV") && m["doubleArrV"] != null ? (m["doubleArrV"] as List<dynamic>).map((a) => a as double).toList(): [],
      m.containsKey("doubleV") && m["doubleV"] != null ? m["doubleV"]: 0.0,
      m.containsKey("intArrV") && m["intArrV"] != null ? (m["intArrV"] as List<dynamic>).map((a) => a as int).toList(): [],
      m.containsKey("intV") && m["intV"] != null ? m["intV"]: 0,
      m.containsKey("objArrV") && m["objArrV"] != null ? (m["objArrV"] as List<dynamic>).map((a) => a as Map<String, dynamic>).map((a) => Json2DartObjArrV.fromMap(a)).toList(): [],
      m.containsKey("objV") && m["objV"] != null ? m["objV"]: null,
      m.containsKey("stringArrV") && m["stringArrV"] != null ? (m["stringArrV"] as List<dynamic>).map((a) => a as String).toList(): [],
      m.containsKey("stringV") && m["stringV"] != null ? m["stringV"]: "",
    );
  }

  Map<String, dynamic> toMap() {
    return {
      "boolArrV": boolArrV,
      "boolV": boolV,
      "doubleArrV": doubleArrV,
      "doubleV": doubleV,
      "intArrV": intArrV,
      "intV": intV,
      "objArrV": objArrV,
      "objV": objV,
      "stringArrV": stringArrV,
      "stringV": stringV,
    };
  }

  String toJson() {
    return '{"boolArrV":$boolArrV, "boolV":$boolV, "doubleArrV":$doubleArrV, "doubleV":$doubleV, "intArrV":$intArrV, "intV":$intV, "objArrV":$objArrV, "objV":$objV, "stringArrV":$stringArrV, "stringV":"$stringV"}';
  }

  @override
  String toString() {
    return '{boolArrV: $boolArrV, boolV: $boolV, doubleArrV: $doubleArrV, doubleV: $doubleV, intArrV: $intArrV, intV: $intV, objArrV: $objArrV, objV: $objV, stringArrV: $stringArrV, stringV: $stringV}';
  }
}

class Json2DartObjArrV {
  final List<bool> boolArrV;
  final bool boolV;
  final List<double> doubleArrV;
  final double doubleV;
  final List<int> intArrV;
  final int intV;
  final List<String> stringArrV;
  final String stringV;

  const Json2DartObjArrV(this.boolArrV, this.boolV, this.doubleArrV, this.doubleV, this.intArrV, this.intV, this.stringArrV, this.stringV);

  factory Json2DartObjArrV.fromMap(Map<String, dynamic> m) {
    return Json2DartObjArrV(
      m.containsKey("boolArrV") && m["boolArrV"] != null ? (m["boolArrV"] as List<dynamic>).map((a) => a as bool).toList(): [],
      m.containsKey("boolV") && m["boolV"] != null ? m["boolV"]: false,
      m.containsKey("doubleArrV") && m["doubleArrV"] != null ? (m["doubleArrV"] as List<dynamic>).map((a) => a as double).toList(): [],
      m.containsKey("doubleV") && m["doubleV"] != null ? m["doubleV"]: 0.0,
      m.containsKey("intArrV") && m["intArrV"] != null ? (m["intArrV"] as List<dynamic>).map((a) => a as int).toList(): [],
      m.containsKey("intV") && m["intV"] != null ? m["intV"]: 0,
      m.containsKey("stringArrV") && m["stringArrV"] != null ? (m["stringArrV"] as List<dynamic>).map((a) => a as String).toList(): [],
      m.containsKey("stringV") && m["stringV"] != null ? m["stringV"]: "",
    );
  }

  Map<String, dynamic> toMap() {
    return {
      "boolArrV": boolArrV,
      "boolV": boolV,
      "doubleArrV": doubleArrV,
      "doubleV": doubleV,
      "intArrV": intArrV,
      "intV": intV,
      "stringArrV": stringArrV,
      "stringV": stringV,
    };
  }

  String toJson() {
    return '{"boolArrV":$boolArrV, "boolV":$boolV, "doubleArrV":$doubleArrV, "doubleV":$doubleV, "intArrV":$intArrV, "intV":$intV, "stringArrV":$stringArrV, "stringV":"$stringV"}';
  }

  @override
  String toString() {
    return '{boolArrV: $boolArrV, boolV: $boolV, doubleArrV: $doubleArrV, doubleV: $doubleV, intArrV: $intArrV, intV: $intV, stringArrV: $stringArrV, stringV: $stringV}';
  }
}

class Json2DartObjV {
  final List<bool> boolArrV;
  final bool boolV;
  final List<double> doubleArrV;
  final double doubleV;
  final List<int> intArrV;
  final int intV;
  final List<String> stringArrV;
  final String stringV;

  const Json2DartObjV(this.boolArrV, this.boolV, this.doubleArrV, this.doubleV, this.intArrV, this.intV, this.stringArrV, this.stringV);

  factory Json2DartObjV.fromMap(Map<String, dynamic> m) {
    return Json2DartObjV(
      m.containsKey("boolArrV") && m["boolArrV"] != null ? (m["boolArrV"] as List<dynamic>).map((a) => a as bool).toList(): [],
      m.containsKey("boolV") && m["boolV"] != null ? m["boolV"]: false,
      m.containsKey("doubleArrV") && m["doubleArrV"] != null ? (m["doubleArrV"] as List<dynamic>).map((a) => a as double).toList(): [],
      m.containsKey("doubleV") && m["doubleV"] != null ? m["doubleV"]: 0.0,
      m.containsKey("intArrV") && m["intArrV"] != null ? (m["intArrV"] as List<dynamic>).map((a) => a as int).toList(): [],
      m.containsKey("intV") && m["intV"] != null ? m["intV"]: 0,
      m.containsKey("stringArrV") && m["stringArrV"] != null ? (m["stringArrV"] as List<dynamic>).map((a) => a as String).toList(): [],
      m.containsKey("stringV") && m["stringV"] != null ? m["stringV"]: "",
    );
  }

  Map<String, dynamic> toMap() {
    return {
      "boolArrV": boolArrV,
      "boolV": boolV,
      "doubleArrV": doubleArrV,
      "doubleV": doubleV,
      "intArrV": intArrV,
      "intV": intV,
      "stringArrV": stringArrV,
      "stringV": stringV,
    };
  }

  String toJson() {
    return '{"boolArrV":$boolArrV, "boolV":$boolV, "doubleArrV":$doubleArrV, "doubleV":$doubleV, "intArrV":$intArrV, "intV":$intV, "stringArrV":$stringArrV, "stringV":"$stringV"}';
  }

  @override
  String toString() {
    return '{boolArrV: $boolArrV, boolV: $boolV, doubleArrV: $doubleArrV, doubleV: $doubleV, intArrV: $intArrV, intV: $intV, stringArrV: $stringArrV, stringV: $stringV}';
  }
}
```
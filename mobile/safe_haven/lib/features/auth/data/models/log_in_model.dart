import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';

class LogInModel extends LogInEntity {
  LogInModel({
    super.fullName,
    required super.password,
    super.phoneNumber,
    required super.email,
  }) {}
  static LogInModel toModel(LogInEntity login_entity) {
    return LogInModel(
      email: login_entity.email,
      password: login_entity.password,
    );
  }

  // From JSON Constructor
  factory LogInModel.fromJson(Map<String, dynamic> json) {
    return LogInModel(
      fullName: json['fullName'],
      password: json['password'],
      phoneNumber: json['phoneNumber'],
      email: json['email'],
    );
  }

  // To JSON Method
  Map<String, dynamic> toJson() {
    return {
      'fullName': fullName ?? '',
      'password': password,
      'phoneNumber': phoneNumber,
      'email': email,
    };
  }

  @override
  String toString() {
    return 'LogInModel( fullName: $fullName, phoneNumber: $phoneNumber, email: $email, password: $password)';
  }
}

class LoggedInModel extends LoggedInEntity {
  LoggedInModel(
      {required super.category,
      required super.language,
      required super.password,
      required super.email,
      super.phoneNumber,
      super.fullName});

  factory LoggedInModel.fromjson(Map<String, dynamic> json) {
    return LoggedInModel(
      category: json['category'],
      language: json['language'],
      password: json['password'],
      fullName: json['fullName'] ?? '',
      phoneNumber: json['phoneNumber'] ?? '',
      email: json['email'] ?? '',
    );
  }
}

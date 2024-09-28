import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';

class LogInModel extends LogInEntity {
  LogInModel({
    required super.userType,
    super.fullName,
    required super.password,
    super.anonymousDifferentiator,
    super.phoneNumber,
    required super.email,
  }) {
    if (userType == 'normal') {
      assert(password.isNotEmpty && email.isNotEmpty,
          'Normal user must have a password and at least one of fullName, phone number, or email.');
    } else if (userType == 'anonymous') {
      assert(anonymousDifferentiator != null && password.isNotEmpty,
          'Anonymous user must have a anonymousDifferentiator and a password.');
    } else {
      throw ArgumentError(
          'Invalid userType: must be either "normal" or "anonymous".');
    }
  }
  static LogInModel toModel(LogInEntity login_entity) {
    return LogInModel(
        email: login_entity.email,
        password: login_entity.password,
        userType: login_entity.userType);
  }

  // From JSON Constructor
  factory LogInModel.fromJson(Map<String, dynamic> json) {
    return LogInModel(
      userType: json['userType'],
      fullName: json['fullName'],
      password: json['password'],
      anonymousDifferentiator: json['anonymousDifferentiator'],
      phoneNumber: json['phoneNumber'],
      email: json['email'],
    );
  }

  // To JSON Method
  Map<String, dynamic> toJson() {
    return {
      'userType': userType,
      'fullName': fullName ?? '',
      'password': password,
      'phoneNumber': phoneNumber,
      'email': email,
      'anonymousDifferentiator': anonymousDifferentiator ?? '',
    };
  }

  // Optional: Validation function
  bool isValid() {
    if (userType == 'normal') {
      return password.isNotEmpty &&
          email.isNotEmpty &&
          (fullName != null || phoneNumber != null);
    } else if (userType == 'anonymous') {
      return password.isNotEmpty && anonymousDifferentiator != null;
    }
    return false;
  }

  @override
  String toString() {
    return 'LogInModel(userType: $userType, fullName: $fullName, anonymousDifferentiator: $anonymousDifferentiator, phoneNumber: $phoneNumber, email: $email, password: $password)';
  }
}

class LoggedInModel extends LoggedInEntity {
  LoggedInModel(
      {required super.userType,
      required super.category,
      required super.language,
      required super.password,
      super.anonymousDifferentiator,
      required super.email,
      super.phoneNumber,
      super.fullName});

  factory LoggedInModel.fromjson(Map<String, dynamic> json) {
    return LoggedInModel(
        userType: json['userType'],
        category: json['category'],
        language: json['language'],
        password: json['password'],
        fullName: json['fullName'] ?? '',
        phoneNumber: json['phoneNumber'] ?? '',
        email: json['email'] ?? '',
        anonymousDifferentiator: json['anonymousDifferentiator'] ?? '');
  }
}

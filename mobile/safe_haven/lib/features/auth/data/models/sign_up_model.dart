import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';

class SignUpModel extends SignUpEntity {
  SignUpModel(
      {required super.userType,
      required super.category,
      required super.language,
      super.anonymousDifferentiator,
      super.email,
      super.phoneNumber,
      required super.fullName,
      required super.password}) {
    if (userType == 'normal') {
      assert(
          password.isNotEmpty &&
              category.isNotEmpty &&
              language.isNotEmpty &&
              fullName.isNotEmpty &&
              (phoneNumber != null || email != null),
          'Normal user must have a password, and either a fullName, an email or a phone number.');
    } else if (userType == 'anonymous') {
      assert(
          password.isNotEmpty &&
              category.isNotEmpty &&
              language.isNotEmpty &&
              anonymousDifferentiator != null,
          'an anonymous user must have anonymousDifferentiator');
    }
  }

  static SignUpModel toModel(SignUpEntity signUpEntity) {
    return SignUpModel(
        userType: signUpEntity.userType,
        category: signUpEntity.category,
        language: signUpEntity.language,
        password: signUpEntity.password,
        email: signUpEntity.email,
        phoneNumber: signUpEntity.phoneNumber,
        fullName: signUpEntity.fullName);
  }

  factory SignUpModel.fromJson(Map<String, dynamic> json) {
    return SignUpModel(
        userType: json['userType'],
        password: json['password'],
        anonymousDifferentiator: json['anonymousDifferentiator'],
        email: json['email'],
        phoneNumber: json['phoneNumber'],
        fullName: json['fullName'],
        language: json['language'],
        category: json['category']);
  }

  // To JSON Method
  Map<String, dynamic> toJson() {
    return {
      'userType': userType,
      'fullName': fullName ,
      'anonymousDifferentiator': anonymousDifferentiator ?? '',
      'email': email ?? '',
      'phoneNumber': phoneNumber ?? '',
      'password': password,
      'language': language,
      'category': category
    };
  }
}

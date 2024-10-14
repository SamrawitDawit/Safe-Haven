import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';

class SignUpModel extends SignUpEntity {
  SignUpModel(
      {
      required super.category,
      required super.language,
      super.email,
      super.phoneNumber,
      required super.fullName,
      required super.password}) {
  
      assert(
          password.isNotEmpty &&
              category.isNotEmpty &&
              language.isNotEmpty &&
              fullName.isNotEmpty &&
              (phoneNumber != null || email != null),
          'Normal user must have a password, and either a fullName, an email or a phone number.');
   
  }

  static SignUpModel toModel(SignUpEntity signUpEntity) {
    return SignUpModel(
        category: signUpEntity.category,
        language: signUpEntity.language,
        password: signUpEntity.password,
        email: signUpEntity.email,
        phoneNumber: signUpEntity.phoneNumber,
        fullName: signUpEntity.fullName);
  }

  factory SignUpModel.fromJson(Map<String, dynamic> json) {
    return SignUpModel(
        password: json['password'],
        email: json['email'],
        phoneNumber: json['phoneNumber'],
        fullName: json['fullName'],
        language: json['language'],
        category: json['category']);
  }

  // To JSON Method
  Map<String, dynamic> toJson() {
    return {
      'fullName': fullName ,
      'email': email ?? '',
      'phoneNumber': phoneNumber ?? '',
      'password': password,
      'language': language,
      'category': category
    };
  }
}

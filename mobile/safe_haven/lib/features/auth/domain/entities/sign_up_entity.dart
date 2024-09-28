import 'package:equatable/equatable.dart';

class SignUpEntity extends Equatable {
  final String userType; // Either 'normal' or 'anonymous'
  final String fullName;
  final String? anonymousDifferentiator;
  final String? email;
  final String? phoneNumber;
  final String password;
  final String language;
  final String category;

  SignUpEntity({
    required this.language,
    required this.category,
    required this.userType,
    required this.fullName,
    this.email,
    this.phoneNumber,
    required this.password,
    this.anonymousDifferentiator,
  }) : assert(
          (userType == 'normal' &&
                  category.isNotEmpty &&
                  language.isNotEmpty &&
                  password.isNotEmpty &&
                  fullName.isNotEmpty &&
                  (email != null || phoneNumber != null)) ||
              (userType == 'anonymous' &&
                  category.isNotEmpty &&
                  language.isNotEmpty &&
                  password.isNotEmpty &&
                  anonymousDifferentiator != null &&
                  fullName.isNotEmpty &&
                  email == null &&
                  phoneNumber == null),
          'Invalid input for user type in sign up entity full name: $fullName, language: $language, password: $password, category: $category, email: $email, phoneNumber : $phoneNumber',
        );

  @override
  List<Object?> get props => [
        fullName,
        email,
        password,
        phoneNumber,
        anonymousDifferentiator,
        userType,
        language,
        category
      ];
}

// class AnonymousSignUpEntity extends Equatable {
//   final String differentiator;
//   final String password;

//   const AnonymousSignUpEntity(
//       {required this.differentiator, required this.password});

//   @override
//   List<Object?> get props => [differentiator, password];
// }

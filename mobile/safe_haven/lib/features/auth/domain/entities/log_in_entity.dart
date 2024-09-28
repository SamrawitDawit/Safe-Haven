import 'package:equatable/equatable.dart';

class LogInEntity extends Equatable {
  final String userType;
  final String? fullName;
  final String password;
  final String email;
  final String? phoneNumber;
  final String? anonymousDifferentiator;

  LogInEntity({
    required this.userType,
    this.fullName,
    required this.password,
    required this.email,
    this.phoneNumber,
    this.anonymousDifferentiator,
  }) : assert(
            (userType == 'normal' && password.isNotEmpty && email.isNotEmpty) ||
                (userType == 'anonymous' &&
                    anonymousDifferentiator != null &&
                    password.isNotEmpty),
            'Invalid input: userType: $userType, fullName = $fullName, phoneNumber = $phoneNumber, email = $email');

  @override
  List<Object?> get props => [
        userType,
        fullName,
        password,
        anonymousDifferentiator,
        phoneNumber,
        email
      ];
}

class LoggedInEntity extends Equatable {
  final String userType;
  final String? fullName;
  final String password;
  final String? anonymousDifferentiator;
  final String? phoneNumber;
  final String email;
  final String language;
  final String category;

  LoggedInEntity(
      {required this.language,
      required this.category,
      required this.userType,
      this.fullName,
      required this.password,
      this.anonymousDifferentiator,
      this.phoneNumber,
      required this.email})
      : assert(
            (userType == 'normal' &&
                    language.isNotEmpty &&
                    category.isNotEmpty &&
                    password.isNotEmpty &&
                    email.isNotEmpty &&
                    (fullName != null || phoneNumber != null)) ||
                (userType == 'anonymous' &&
                    anonymousDifferentiator != null &&
                    password.isNotEmpty),
            'invalid input for login lorfreesfd user type');

  @override
  List<Object?> get props => [
        userType,
        fullName,
        password,
        anonymousDifferentiator,
        phoneNumber,
        email,
        language,
        category
      ];
}

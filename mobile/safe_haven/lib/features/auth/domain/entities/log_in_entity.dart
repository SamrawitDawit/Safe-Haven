import 'package:equatable/equatable.dart';

class LogInEntity extends Equatable {
  final String? fullName;
  final String password;
  final String email;
  final String? phoneNumber;

  LogInEntity({
    this.fullName,
    required this.password,
    required this.email,
    this.phoneNumber,
  }) : assert(
            (password.isNotEmpty && email.isNotEmpty),
            'Invalid input: fullName = $fullName, phoneNumber = $phoneNumber, email = $email');

  @override
  List<Object?> get props => [
        fullName,
        password,
        phoneNumber,
        email
      ];
}

class LoggedInEntity extends Equatable {
  final String? fullName;
  final String password;
  final String? phoneNumber;
  final String email;
  final String language;
  final String category;

  LoggedInEntity(
      {required this.language,
      required this.category,
      this.fullName,
      required this.password,
      this.phoneNumber,
      required this.email})
      : assert(
            (
                    language.isNotEmpty &&
                    category.isNotEmpty &&
                    password.isNotEmpty &&
                    email.isNotEmpty &&
                    (fullName != null || phoneNumber != null)),
            'invalid input for login in entities');

  @override
  List<Object?> get props => [
        fullName,
        password,
        phoneNumber,
        email,
        language,
        category
      ];
}

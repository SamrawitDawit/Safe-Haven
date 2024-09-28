import 'package:dartz/dartz.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';
import 'package:safe_haven/features/auth/domain/entities/reset_password_entity.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';

abstract class AuthenticationRepository {
  Future<Either<Failure, Unit>> logIn(LogInEntity logInEntity);
  Future<Either<Failure, Unit>> signUp(SignUpEntity signUpEntity);
  Future<Either<Failure, Unit>> logOut();
  Future<Either<Failure, Unit>> forgotPassword(String resetEmail);
  Future<Either<Failure, Unit>> resetPassword(
      ResetPasswordEntity resetPasswordEntity);
  Future<Either<Failure, Unit>> googleSignIn();
}

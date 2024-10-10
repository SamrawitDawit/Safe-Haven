part of 'auth_bloc_bloc.dart';

@immutable
sealed class AuthBlocState extends Equatable {}

final class AuthBlocInitial extends AuthBlocState {
  @override
  List<Object?> get props => [];
}

class AuthLoading extends AuthBlocState {
  @override
  List<Object?> get props => [];
}

class AuthError extends AuthBlocState {
  final String message;

  AuthError({required this.message});

  @override
  List<Object> get props => [message];
}

class LoggedOut extends AuthBlocState {
  @override
  List<Object?> get props => [];
}

class LogInSuccess extends AuthBlocState {
  final String logInSuccessMessage;

  LogInSuccess({required this.logInSuccessMessage});
  @override
  List<Object?> get props => [logInSuccessMessage];
}

class LoggInError extends AuthBlocState {
  final String logInErrorMessage;

  LoggInError({required this.logInErrorMessage});
  @override
  List<Object?> get props => [logInErrorMessage];
}

class AuthSuccess extends AuthBlocState {
  final String successMessage;

  AuthSuccess({required this.successMessage});
  @override
  List<Object?> get props => [successMessage];
}

class AuthRegisterSuccess extends AuthBlocState {
  final String successMessage;

  AuthRegisterSuccess({required this.successMessage});
  @override
  List<Object?> get props => [successMessage];
}

class ForgotPasswordSuccess extends AuthBlocState {
  final String forgotpasswordsuccess;

  ForgotPasswordSuccess({required this.forgotpasswordsuccess});
  @override
  List<Object?> get props => [forgotpasswordsuccess];
}

class ForgotPasswordError extends AuthBlocState {
  final String forgotPasswordErrorMessage;

  ForgotPasswordError({required this.forgotPasswordErrorMessage});
  @override
  List<Object?> get props => [forgotPasswordErrorMessage];
}

class ForgotLoading extends AuthBlocState {
  @override
  List<Object?> get props => [];
}

class ResetPasswordLoading extends AuthBlocState {
  @override
  List<Object?> get props => [];
}

class ResetPasswordSuccessState extends AuthBlocState {
  final String successResetPasswordMessage;

  ResetPasswordSuccessState({required this.successResetPasswordMessage});
  @override
  List<Object?> get props => [successResetPasswordMessage];
}

class ResetPasswordErrorState extends AuthBlocState {
  final String errorResetPasswordMessage;

  ResetPasswordErrorState({required this.errorResetPasswordMessage});
  @override
  List<Object?> get props => [errorResetPasswordMessage];
}

class GoogleSigninState extends AuthBlocState {
  final String successGoogleSignInMessage;

  GoogleSigninState({required this.successGoogleSignInMessage});

  @override
  List<Object?> get props => [successGoogleSignInMessage];
}

class GoogleSignInLoading extends AuthBlocState {
  @override
  List<Object?> get props => [];
}

class GoogleSignInErrorState extends AuthBlocState {
  final String errorGoogleSignInMessage;

  GoogleSignInErrorState({required this.errorGoogleSignInMessage});
  @override
  List<Object?> get props => [errorGoogleSignInMessage];
}

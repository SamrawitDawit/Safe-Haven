part of 'auth_bloc_bloc.dart';

@immutable
sealed class AuthBlocEvent extends Equatable {
  const AuthBlocEvent();

  @override
  List<Object> get props => [];
}

class LoginEvent extends AuthBlocEvent {
  final LogInEntity loginEntity;

  const LoginEvent({required this.loginEntity});
}

class RegisterEvent extends AuthBlocEvent {
  final SignUpEntity registrationEntity;

  const RegisterEvent({required this.registrationEntity});
}

class LogoutEvent extends AuthBlocEvent {}

class GetUserEvent extends AuthBlocEvent {}

class ForgotPasswordEvent extends AuthBlocEvent {
  final String resetEmail;

  const ForgotPasswordEvent({required this.resetEmail});
}

class ResetPasswordEvent extends AuthBlocEvent {
  final ResetPasswordEntity resetPasswordEntity;

  const ResetPasswordEvent({required this.resetPasswordEntity});
}

class GoogleSignInEvent extends AuthBlocEvent {
  const GoogleSignInEvent();
}

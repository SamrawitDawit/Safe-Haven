import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:meta/meta.dart';
import 'package:safe_haven/core/BaseUseCase/base_usecase.dart';
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';
import 'package:safe_haven/features/auth/domain/entities/reset_password_entity.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';
import 'package:safe_haven/features/auth/domain/usecases/forgot_password.dart';
import 'package:safe_haven/features/auth/domain/usecases/google_sign_in.dart';
import 'package:safe_haven/features/auth/domain/usecases/log_in.dart';
import 'package:safe_haven/features/auth/domain/usecases/log_out.dart';
import 'package:safe_haven/features/auth/domain/usecases/reset_password.dart';
import 'package:safe_haven/features/auth/domain/usecases/sign_up.dart';

part 'auth_bloc_event.dart';
part 'auth_bloc_state.dart';

class AuthBlocBloc extends Bloc<AuthBlocEvent, AuthBlocState> {
  final SignUpUsecase signUpUsecase;
  final LogInUsecase logInUsecase;
  final LogoutUsecase logoutUsecase;
  final ForgotPasswordUsecase forgotPasswordUsecase;
  final ResetPasswordUseCase resetPasswordUseCase;
  final GoogleSignInUseCase googleSignInUseCase;
  AuthBlocBloc(
      this.signUpUsecase,
      this.logInUsecase,
      this.logoutUsecase,
      this.forgotPasswordUsecase,
      this.resetPasswordUseCase,
      this.googleSignInUseCase)
      : super(AuthBlocInitial()) {
    on<RegisterEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await signUpUsecase(
          SignUpParams(signUpEntity: event.registrationEntity));

      result.fold((failure) {
        print(failure);
        emit(AuthError(message: failure.errorMessage));
      }, (data) {
        emit(AuthRegisterSuccess(
            successMessage: 'successfully registered you!'));
      });
    });
    on<LoginEvent>((event, emit) async {
      emit(AuthLoading());
      final result =
          await logInUsecase(LogInParams(loginEntity: event.loginEntity));

      result.fold((failure) {
        emit(AuthError(message: 'failed in being logged in '));
      }, (data) {
        emit(LogInSuccess(logInSuccessMessage: 'successfully logged u in'));
      });
    });
    on<LogoutEvent>((event, emit) async {
      emit(AuthLoading());
      final result = await logoutUsecase(NoParams());

      result.fold((failure) {
        emit(AuthError(message: 'failed in being logged out'));
      }, (data) {
        emit(
            AuthRegisterSuccess(successMessage: 'successfully Logged you in!'));
      });
    });

    on<ForgotPasswordEvent>((event, emit) async {
      emit(ForgotLoading());
      final result = await forgotPasswordUsecase(
          ForgotPasswordParams(resetEmail: event.resetEmail));
      result.fold((failure) {
        emit(ForgotPasswordError(
            forgotPasswordErrorMessage:
                'error in bloc of sending reset email'));
      }, (data) {
        emit(ForgotPasswordSuccess(
            forgotpasswordsuccess: 'successfuly sent reset password email'));
      });
    });

    on<ResetPasswordEvent>((event, emit) async {
      emit(ResetPasswordLoading());
      final result = await resetPasswordUseCase(
          ResetPasswordParams(resetPasswordEntity: event.resetPasswordEntity));
      result.fold((failure) {
        emit(ResetPasswordErrorState(
            errorResetPasswordMessage:
                'error in resetting password in the bloc '));
      }, (data) {
        emit(ResetPasswordSuccessState(
            successResetPasswordMessage:
                'successfully reset password in the bloc'));
      });
    });

    on<GoogleSignInEvent>((event, emit) async {
      emit(GoogleSignInLoading());
      final result = await googleSignInUseCase(NoParams());
      result.fold((failure) {
        emit(GoogleSignInErrorState(
            errorGoogleSignInMessage: 'error when signing in with google'));
      }, (data) {
        emit(AuthSuccess(successMessage: 'successfully signed it with google'));
      });
    });
  }
}

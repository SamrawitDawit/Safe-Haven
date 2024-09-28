import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
import 'package:safe_haven/core/BaseUseCase/base_usecase.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';

class ForgotPasswordUsecase extends BaseUsecase<Unit, ForgotPasswordParams> {
  final AuthenticationRepository authenticationRepository;

  ForgotPasswordUsecase({required this.authenticationRepository});

  Future<Either<Failure, Unit>> call(
      ForgotPasswordParams forgotPasswordParams) async {
    print('teyertao;');
    return await authenticationRepository
        .forgotPassword(forgotPasswordParams.resetEmail);
  }
}

class ForgotPasswordParams extends Equatable {
  final String resetEmail;

  ForgotPasswordParams({required this.resetEmail});

  @override
  List<Object?> get props => [resetEmail];
}

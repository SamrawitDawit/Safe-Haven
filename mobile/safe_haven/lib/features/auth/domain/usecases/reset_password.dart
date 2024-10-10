import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
import 'package:safe_haven/core/BaseUseCase/base_usecase.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/domain/entities/reset_password_entity.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';

class ResetPasswordUseCase extends BaseUsecase<Unit, ResetPasswordParams> {
  final AuthenticationRepository authenticationRepository;

  ResetPasswordUseCase({required this.authenticationRepository});

  Future<Either<Failure, Unit>> call(
      ResetPasswordParams ResetPasswordParams) async {
    print('teyertao;');
    return await authenticationRepository
        .resetPassword(ResetPasswordParams.resetPasswordEntity);
  }
}

class ResetPasswordParams extends Equatable {
  final ResetPasswordEntity resetPasswordEntity;

  ResetPasswordParams({required this.resetPasswordEntity});

  @override
  List<Object?> get props => [resetPasswordEntity];
}

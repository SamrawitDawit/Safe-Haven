import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
import 'package:safe_haven/core/BaseUseCase/base_usecase.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';

class LogInUsecase extends BaseUsecase<Unit, LogInParams> {
  final AuthenticationRepository authenticationRepository;

  LogInUsecase({required this.authenticationRepository});

  Future<Either<Failure, Unit>> call(LogInParams LogInParams) async {
    print('teyertao;');
    return await authenticationRepository.logIn(LogInParams.loginEntity);
  }
}

class LogInParams extends Equatable {
  final LogInEntity loginEntity;

  LogInParams({required this.loginEntity});

  @override
  List<Object?> get props => [loginEntity];
}

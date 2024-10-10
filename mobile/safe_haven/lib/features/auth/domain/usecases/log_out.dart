import 'package:dartz/dartz.dart';
import 'package:safe_haven/core/BaseUseCase/base_usecase.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';

class LogoutUsecase implements BaseUsecase<Unit, NoParams> {
  final AuthenticationRepository authenticationRepository;

LogoutUsecase({required this.authenticationRepository});

  @override
  Future<Either<Failure, Unit>> call(NoParams params) {
    return authenticationRepository.logOut();
  }
}

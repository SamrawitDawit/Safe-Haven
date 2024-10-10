import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
import 'package:safe_haven/core/BaseUseCase/base_usecase.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';

class GoogleSignInUseCase extends BaseUsecase<Unit, NoParams> {
  final AuthenticationRepository authenticationRepository;

  GoogleSignInUseCase({required this.authenticationRepository});
  @override
  Future<Either<Failure, Unit>> call(NoParams params) async {
    return await authenticationRepository.googleSignIn();
  }
}

class GoogleParams extends Equatable {
  const GoogleParams();

  @override
  List<Object?> get props => [];
}

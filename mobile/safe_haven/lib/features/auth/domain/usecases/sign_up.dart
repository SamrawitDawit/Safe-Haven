import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
import 'package:safe_haven/core/BaseUseCase/base_usecase.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';

class SignUpUsecase extends BaseUsecase<Unit, SignUpParams> {
  final AuthenticationRepository authenticationRepository;

  SignUpUsecase({required this.authenticationRepository});

  Future<Either<Failure, Unit>> call(SignUpParams signUpParams) async {
    return await authenticationRepository.signUp(signUpParams.signUpEntity);
  }
}

class SignUpParams extends Equatable {
  final SignUpEntity signUpEntity;

  SignUpParams({required this.signUpEntity});

  @override
  List<Object?> get props => [];
}

import 'package:dartz/dartz.dart';
import 'package:flutter/foundation.dart';
import 'package:safe_haven/core/constants/constants.dart';
import 'package:safe_haven/core/error/exception.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/core/network/network_info.dart';
import 'package:safe_haven/features/auth/data/data_sources/local_data_source.dart';
import 'package:safe_haven/features/auth/data/data_sources/remote_data_source.dart';
import 'package:safe_haven/features/auth/data/models/log_in_model.dart';
import 'package:safe_haven/features/auth/data/models/reset_password_model.dart';
import 'package:safe_haven/features/auth/data/models/sign_up_model.dart';
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';
import 'package:safe_haven/features/auth/domain/entities/reset_password_entity.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';

class AuthenticationRepoImpl implements AuthenticationRepository {
  final AuthenticationRemoteDataSource remoteDataSource;
  final AuthenticationLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  AuthenticationRepoImpl(
      {required this.remoteDataSource,
      required this.localDataSource,
      required this.networkInfo});
  @override
  Future<Either<Failure, Unit>> logIn(LogInEntity logInEntity) async {
    if (await networkInfo.isConnected) {
      try {
        print('hello');
        final result =
            await remoteDataSource.login(LogInModel.toModel(logInEntity));

        try {
          print(result.accessToken);
          await localDataSource.cacheTokens(
              result.accessToken, result.refreshToken);
          await localDataSource.storeUser(result);
        } on CacheException {
          debugPrint('Caching Token Error');
        }
        return const Right(unit);
      } on UnauthorizedException {
        return const Left(UnauthorizedFailure('unauthorized'));
        // ignore: dead_code_on_catch_subtype
      } on ServerException {
        return const Left(ServerException('server exception'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, Unit>> logOut() async {
    try {
      await localDataSource.logout();
      return const Right(unit);
    } on CacheException {
      return const Left(CacheFailure(ErrorMessages.cacheError));
    }
  }

  @override
  Future<Either<Failure, Unit>> signUp(SignUpEntity signUpEntity) async {
    networkInfo.isConnected;
    if (await networkInfo.isConnected) {
      try {
        print('hello');
        final result =
            await remoteDataSource.signup(SignUpModel.toModel(signUpEntity));
        return Right(result);
      } on UnauthorizedException {
        return const Left(UnauthorizedFailure('unauthorized'));
        // ignore: dead_code_on_catch_subtype
      } on ServerException {
        return const Left(ServerException('server exception'));
      } on SocketException {
        return const Left(ConnectionFailure('Socket connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, Unit>> forgotPassword(String resetEmail) async {
    networkInfo.isConnected;
    if (await networkInfo.isConnected) {
      try {
        print('hello');
        final result = await remoteDataSource.forgotPassword(resetEmail);
        return Right(result);
      } on UnauthorizedException {
        return const Left(UnauthorizedFailure('unauthorized'));
        // ignore: dead_code_on_catch_subtype
      } on ServerException {
        return const Left(ServerException('server exception'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, Unit>> resetPassword(
      ResetPasswordEntity resetPasswordEntity) async {
    networkInfo.isConnected;
    if (await networkInfo.isConnected) {
      try {
        print('hello');
        final result = await remoteDataSource
            .resetPassword(ResetPasswordModel.toModel(resetPasswordEntity));
        return Right(result);
      } on UnauthorizedException {
        return const Left(UnauthorizedFailure('unauthorized'));
        // ignore: dead_code_on_catch_subtype
      } on ServerException {
        return const Left(ServerException('server exception'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }

  @override
  Future<Either<Failure, Unit>> googleSignIn() async {
    networkInfo.isConnected;
    if (await networkInfo.isConnected) {
      try {
        print('hello');
        await remoteDataSource.googleLogin();
        return const Right(unit);
      } on UnauthorizedException {
        return const Left(UnauthorizedFailure('unauthorized'));
      } on ServerException {
        return const Left(ServerException('server exception'));
      } on SocketException {
        return const Left(ConnectionFailure('No internet connection'));
      }
    } else {
      return const Left(ConnectionFailure('No internet connection'));
    }
  }
}

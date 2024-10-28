import 'package:dartz/dartz.dart';
import 'package:safe_haven/core/error/exception.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/case/data/data_sources/local_data_source_case.dart';
import 'package:safe_haven/features/case/data/data_sources/remote_data_source_case.dart';
import 'package:safe_haven/features/case/data/models/case_model.dart';
import 'package:safe_haven/features/case/domain/entities/case_entity.dart';
import 'package:safe_haven/features/case/domain/repositories/case_repository.dart';
import 'package:safe_haven/core/network/network_info.dart';

class CaseRepoImpl implements CaseRepository {
  final CaseRemoteDataSource caseRemoteDataSource;
  final CaseLocalDataSource caseLocalDataSource;
  final NetworkInfo networkInfo;

  CaseRepoImpl(
      {required this.caseRemoteDataSource,
      required this.caseLocalDataSource,
      required this.networkInfo});

  @override
  Future<Either<Failure, CaseEntity>> createCase(CaseEntity caseEntity) async {
    networkInfo.isConnected;
    if (await networkInfo.isConnected) {
      try {
        print('hello');
        final result = await caseRemoteDataSource
            .createCase(CaseModel.toModel(caseEntity));
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
  Future<Either<Failure, Unit>> deleteCase(String id) {
    // TODO: implement deleteCase
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, CaseEntity>> getCase(CaseEntity caseEntity) {
    // TODO: implement getCase
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, List<CaseEntity>>> getCases() {
    // TODO: implement getCases
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, CaseEntity>> updateCase(String id) {
    // TODO: implement updateCase
    throw UnimplementedError();
  }
}

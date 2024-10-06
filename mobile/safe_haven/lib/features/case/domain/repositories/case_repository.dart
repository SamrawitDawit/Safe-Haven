import 'package:dartz/dartz.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/case/domain/entities/case_entity.dart';

abstract class CaseRepository {
  Future<Either<Failure, CaseEntity>> createCase(CaseEntity caseEntity);
  Future<Either<Failure, CaseEntity>> updateCase(String id);
  Future<Either<Failure, Unit>> deleteCase(String id);
  Future<Either<Failure, CaseEntity>> getCase(CaseEntity caseEntity);
  Future<Either<Failure, List<CaseEntity>>> getCases();
}

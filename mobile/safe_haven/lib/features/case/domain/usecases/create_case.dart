import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
import 'package:safe_haven/core/BaseUseCase/base_usecase.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/case/domain/entities/case_entity.dart';
import 'package:safe_haven/features/case/domain/repositories/case_repository.dart';

class CreateCaseUseCase extends BaseUsecase<CaseEntity, CreateCaseParams> {
  final CaseRepository caseRepository;

  CreateCaseUseCase({required this.caseRepository});

  @override
  Future<Either<Failure, CaseEntity>> call(CreateCaseParams params) async {
    return await caseRepository.createCase(params.caseEntity);
  }
}

class CreateCaseParams extends Equatable {
  final CaseEntity caseEntity;

  CreateCaseParams({required this.caseEntity});

  @override
  List<Object?> get props => [caseEntity];
}

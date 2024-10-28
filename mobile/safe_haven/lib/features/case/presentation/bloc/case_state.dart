part of 'case_bloc.dart';

@immutable
abstract class CaseState extends Equatable {
  const CaseState();
}

final class CaseInitial extends CaseState {
  @override
  List<Object?> get props => [];
}

class LoadedSingleCaseState extends CaseState {
  final CaseEntity singleCase;

  const LoadedSingleCaseState(this.singleCase);

  @override
  List<Object?> get props => [singleCase];
}

class CaseLoadError extends CaseState {
  final String message;

  const CaseLoadError(this.message);

  @override
  List<Object?> get props => [message];
}

//get all Cases

class LoadedAllCases extends CaseState {
  final List<CaseEntity> Cases;
  const LoadedAllCases(this.Cases);

  @override
  List<Object?> get props => [Cases];
}

class LoadingAllCase extends CaseState {
  const LoadingAllCase();
  @override
  List<Object?> get props => [];
}

class LoadingAllCasesError extends CaseState {
  final String message;

  const LoadingAllCasesError(this.message);

  @override
  List<Object?> get props => [message];
}

class UpdatedCase extends CaseState {
  final CaseEntity Case;

  const UpdatedCase(this.Case);
  @override
  List<Object?> get props => [Case];
}

class UpdatingCase extends CaseState {
  const UpdatingCase();
  @override
  List<Object?> get props => [];
}

class UpdateCaseError extends CaseState {
  final String message;
  const UpdateCaseError(this.message);

  @override
  List<Object?> get props => [message];
}

class CreatedCase extends CaseState {
  final CaseEntity Case;

  const CreatedCase(this.Case);
  @override
  List<Object?> get props => [Case];
}

class CreatingCase extends CaseState {
  const CreatingCase();
  @override
  List<Object?> get props => [];
}

class CreateCaseError extends CaseState {
  final String errorMessage;
  const CreateCaseError({required this.errorMessage});

  @override
  List<Object?> get props => [errorMessage];
}

class DeletedCase extends CaseState {
  const DeletedCase();
  @override
  List<Object?> get props => [];
}

class DeletingCase extends CaseState {
  const DeletingCase();
  @override
  List<Object?> get props => [];
}

class DeleteCaseError extends CaseState {
  final String message;
  const DeleteCaseError(this.message);

  @override
  List<Object?> get props => [message];
}

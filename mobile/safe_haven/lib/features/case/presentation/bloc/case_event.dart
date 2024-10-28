part of 'case_bloc.dart';

@immutable
abstract class CaseEvent extends Equatable {
  const CaseEvent();
}

class GetSingleCaseEvent extends CaseEvent {
  final String id;

  const GetSingleCaseEvent(this.id);

  @override
  List<Object?> get props => [id];
}

class LoadAllCasesEvent extends CaseEvent {
  const LoadAllCasesEvent();

  @override
  List<Object?> get props => [];
}

class UpdateCaseEvent extends CaseEvent {
  final CaseEntity singleCase;
  const UpdateCaseEvent(this.singleCase);

  @override
  List<Object?> get props => throw UnimplementedError();
}

class CreateCaseEvent extends CaseEvent {
  final CaseEntity singleCase;
  const CreateCaseEvent({required this.singleCase});

  @override
  List<Object?> get props => [singleCase];
}

class DeleteCaseEvent extends CaseEvent {
  final String id;

  const DeleteCaseEvent({required this.id});

  @override
  List<Object?> get props => [id];
}

class ResetCaseEvent extends CaseEvent {
  @override
  List<Object?> get props => [];
}
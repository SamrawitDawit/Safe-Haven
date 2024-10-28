import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:meta/meta.dart';
import 'package:safe_haven/features/case/domain/entities/case_entity.dart';
import 'package:safe_haven/features/case/domain/usecases/create_case.dart';

part 'case_event.dart';
part 'case_state.dart';

class CaseBloc extends Bloc<CaseEvent, CaseState> {
  final CreateCaseUseCase createCaseUseCase;

  CaseBloc(this.createCaseUseCase) : super(CaseInitial()) {
    on<CreateCaseEvent>((event, emit) async {
      emit(CreatingCase());
      final result = await createCaseUseCase(
          CreateCaseParams(caseEntity: event.singleCase));
      result.fold((failure) {
        emit(CreateCaseError(errorMessage: failure.errorMessage));
      }, (data) {
        emit(CreatedCase(data));
      });
    });
    on<ResetCaseEvent>(
      (event, emit) => emit(CaseInitial()),
    );
  }
}

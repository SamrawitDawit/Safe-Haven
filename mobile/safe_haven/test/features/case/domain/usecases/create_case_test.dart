import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:safe_haven/features/case/domain/entities/case_entity.dart';
import 'package:safe_haven/features/case/domain/usecases/create_case.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late CreateCaseUseCase createCaseUseCase;
  late MockCaseRepository mockCaseRepository;

  setUp(() {
    mockCaseRepository = MockCaseRepository();
    createCaseUseCase = CreateCaseUseCase(caseRepository: mockCaseRepository);
  });

  final testCase = CaseEntity(
      id: 'id',
      title: 'title',
      description: 'description',
      image_url: 'image_url');

  test('test that case is created and reported ', () async {
    // arrange
    when(mockCaseRepository.createCase(testCase))
        .thenAnswer((_) async => Right(testCase));

    // act
    final result =
        await createCaseUseCase(CreateCaseParams(caseEntity: testCase));

    //assert
    expect(result, Right(testCase));
  });
}

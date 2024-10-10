import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';
import 'package:safe_haven/features/auth/domain/usecases/sign_up.dart';

import '../../../helpers/test_helper.mocks.dart';

void main() {
  late SignUpUsecase signUpUsecase;
  late MockAuthenticationRepository mockAuthenticationRepository;

  setUp(() {
    mockAuthenticationRepository = MockAuthenticationRepository();
    signUpUsecase =
        SignUpUsecase(authenticationRepository: mockAuthenticationRepository);
  });

  final testSignUpEntity = SignUpEntity( password: 'pass', fullName: 'bereket', phoneNumber: '987', language: 'language', category: 'cat');
  test('tests that correct url is called to sign up', () async {
    //arrange
    when(mockAuthenticationRepository.signUp(testSignUpEntity))
    .thenAnswer((_) async => Right(unit) );

    //act
    final result = await signUpUsecase(SignUpParams(signUpEntity: testSignUpEntity));

    //assert
    expect(result, Right(unit));
    verify(mockAuthenticationRepository.signUp(testSignUpEntity));
    verifyNoMoreInteractions(mockAuthenticationRepository);
  });

}

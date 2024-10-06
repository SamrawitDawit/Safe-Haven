// Mocks generated by Mockito 5.4.4 from annotations
// in safe_haven/test/helpers/test_helper.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i7;
import 'dart:convert' as _i22;
import 'dart:typed_data' as _i23;

import 'package:dartz/dartz.dart' as _i2;
import 'package:http/http.dart' as _i5;
import 'package:internet_connection_checker/internet_connection_checker.dart'
    as _i4;
import 'package:mockito/mockito.dart' as _i1;
import 'package:mockito/src/dummies.dart' as _i17;
import 'package:safe_haven/core/error/faliure.dart' as _i8;
import 'package:safe_haven/core/network/network_info.dart' as _i18;
import 'package:safe_haven/features/auth/data/data_sources/local_data_source.dart'
    as _i16;
import 'package:safe_haven/features/auth/data/data_sources/remote_data_source.dart'
    as _i12;
import 'package:safe_haven/features/auth/data/models/authenticated_model.dart'
    as _i3;
import 'package:safe_haven/features/auth/data/models/log_in_model.dart' as _i13;
import 'package:safe_haven/features/auth/data/models/reset_password_model.dart'
    as _i15;
import 'package:safe_haven/features/auth/data/models/sign_up_model.dart'
    as _i14;
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart'
    as _i9;
import 'package:safe_haven/features/auth/domain/entities/reset_password_entity.dart'
    as _i11;
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart'
    as _i10;
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart'
    as _i6;
import 'package:safe_haven/features/case/domain/entities/case_entity.dart'
    as _i21;
import 'package:safe_haven/features/case/domain/repositories/case_repository.dart'
    as _i20;
import 'package:shared_preferences/shared_preferences.dart' as _i19;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: deprecated_member_use
// ignore_for_file: deprecated_member_use_from_same_package
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

class _FakeEither_0<L, R> extends _i1.SmartFake implements _i2.Either<L, R> {
  _FakeEither_0(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeAuthenticatedModel_1 extends _i1.SmartFake
    implements _i3.AuthenticatedModel {
  _FakeAuthenticatedModel_1(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeUnit_2 extends _i1.SmartFake implements _i2.Unit {
  _FakeUnit_2(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeDuration_3 extends _i1.SmartFake implements Duration {
  _FakeDuration_3(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeAddressCheckResult_4 extends _i1.SmartFake
    implements _i4.AddressCheckResult {
  _FakeAddressCheckResult_4(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeResponse_5 extends _i1.SmartFake implements _i5.Response {
  _FakeResponse_5(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeStreamedResponse_6 extends _i1.SmartFake
    implements _i5.StreamedResponse {
  _FakeStreamedResponse_6(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

/// A class which mocks [AuthenticationRepository].
///
/// See the documentation for Mockito's code generation for more information.
class MockAuthenticationRepository extends _i1.Mock
    implements _i6.AuthenticationRepository {
  MockAuthenticationRepository() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>> logIn(
          _i9.LogInEntity? logInEntity) =>
      (super.noSuchMethod(
        Invocation.method(
          #logIn,
          [logInEntity],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>.value(
            _FakeEither_0<_i8.Failure, _i2.Unit>(
          this,
          Invocation.method(
            #logIn,
            [logInEntity],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>> signUp(
          _i10.SignUpEntity? signUpEntity) =>
      (super.noSuchMethod(
        Invocation.method(
          #signUp,
          [signUpEntity],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>.value(
            _FakeEither_0<_i8.Failure, _i2.Unit>(
          this,
          Invocation.method(
            #signUp,
            [signUpEntity],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>> logOut() => (super.noSuchMethod(
        Invocation.method(
          #logOut,
          [],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>.value(
            _FakeEither_0<_i8.Failure, _i2.Unit>(
          this,
          Invocation.method(
            #logOut,
            [],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>> forgotPassword(
          String? resetEmail) =>
      (super.noSuchMethod(
        Invocation.method(
          #forgotPassword,
          [resetEmail],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>.value(
            _FakeEither_0<_i8.Failure, _i2.Unit>(
          this,
          Invocation.method(
            #forgotPassword,
            [resetEmail],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>> resetPassword(
          _i11.ResetPasswordEntity? resetPasswordEntity) =>
      (super.noSuchMethod(
        Invocation.method(
          #resetPassword,
          [resetPasswordEntity],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>.value(
            _FakeEither_0<_i8.Failure, _i2.Unit>(
          this,
          Invocation.method(
            #resetPassword,
            [resetPasswordEntity],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>> googleSignIn() =>
      (super.noSuchMethod(
        Invocation.method(
          #googleSignIn,
          [],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>.value(
            _FakeEither_0<_i8.Failure, _i2.Unit>(
          this,
          Invocation.method(
            #googleSignIn,
            [],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>);
}

/// A class which mocks [AuthenticationRemoteDataSource].
///
/// See the documentation for Mockito's code generation for more information.
class MockAuthenticationRemoteDataSource extends _i1.Mock
    implements _i12.AuthenticationRemoteDataSource {
  MockAuthenticationRemoteDataSource() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i7.Future<_i3.AuthenticatedModel> login(_i13.LogInModel? logInModel) =>
      (super.noSuchMethod(
        Invocation.method(
          #login,
          [logInModel],
        ),
        returnValue:
            _i7.Future<_i3.AuthenticatedModel>.value(_FakeAuthenticatedModel_1(
          this,
          Invocation.method(
            #login,
            [logInModel],
          ),
        )),
      ) as _i7.Future<_i3.AuthenticatedModel>);

  @override
  _i7.Future<_i2.Unit> signup(_i14.SignUpModel? signUpModel) =>
      (super.noSuchMethod(
        Invocation.method(
          #signup,
          [signUpModel],
        ),
        returnValue: _i7.Future<_i2.Unit>.value(_FakeUnit_2(
          this,
          Invocation.method(
            #signup,
            [signUpModel],
          ),
        )),
      ) as _i7.Future<_i2.Unit>);

  @override
  _i7.Future<_i2.Unit> forgotPassword(String? resetEmail) =>
      (super.noSuchMethod(
        Invocation.method(
          #forgotPassword,
          [resetEmail],
        ),
        returnValue: _i7.Future<_i2.Unit>.value(_FakeUnit_2(
          this,
          Invocation.method(
            #forgotPassword,
            [resetEmail],
          ),
        )),
      ) as _i7.Future<_i2.Unit>);

  @override
  _i7.Future<_i2.Unit> resetPassword(
          _i15.ResetPasswordModel? resetPasswordModel) =>
      (super.noSuchMethod(
        Invocation.method(
          #resetPassword,
          [resetPasswordModel],
        ),
        returnValue: _i7.Future<_i2.Unit>.value(_FakeUnit_2(
          this,
          Invocation.method(
            #resetPassword,
            [resetPasswordModel],
          ),
        )),
      ) as _i7.Future<_i2.Unit>);

  @override
  _i7.Future<_i2.Unit> googleLogin() => (super.noSuchMethod(
        Invocation.method(
          #googleLogin,
          [],
        ),
        returnValue: _i7.Future<_i2.Unit>.value(_FakeUnit_2(
          this,
          Invocation.method(
            #googleLogin,
            [],
          ),
        )),
      ) as _i7.Future<_i2.Unit>);
}

/// A class which mocks [AuthenticationLocalDataSource].
///
/// See the documentation for Mockito's code generation for more information.
class MockAuthenticationLocalDataSource extends _i1.Mock
    implements _i16.AuthenticationLocalDataSource {
  MockAuthenticationLocalDataSource() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i7.Future<String> getToken() => (super.noSuchMethod(
        Invocation.method(
          #getToken,
          [],
        ),
        returnValue: _i7.Future<String>.value(_i17.dummyValue<String>(
          this,
          Invocation.method(
            #getToken,
            [],
          ),
        )),
      ) as _i7.Future<String>);

  @override
  _i7.Future<_i2.Unit> cacheTokens(
    String? token,
    String? refreshToken,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #cacheTokens,
          [
            token,
            refreshToken,
          ],
        ),
        returnValue: _i7.Future<_i2.Unit>.value(_FakeUnit_2(
          this,
          Invocation.method(
            #cacheTokens,
            [
              token,
              refreshToken,
            ],
          ),
        )),
      ) as _i7.Future<_i2.Unit>);

  @override
  _i7.Future<_i2.Unit> logout() => (super.noSuchMethod(
        Invocation.method(
          #logout,
          [],
        ),
        returnValue: _i7.Future<_i2.Unit>.value(_FakeUnit_2(
          this,
          Invocation.method(
            #logout,
            [],
          ),
        )),
      ) as _i7.Future<_i2.Unit>);

  @override
  _i7.Future<String> getRefreshToken() => (super.noSuchMethod(
        Invocation.method(
          #getRefreshToken,
          [],
        ),
        returnValue: _i7.Future<String>.value(_i17.dummyValue<String>(
          this,
          Invocation.method(
            #getRefreshToken,
            [],
          ),
        )),
      ) as _i7.Future<String>);
}

/// A class which mocks [NetworkInfo].
///
/// See the documentation for Mockito's code generation for more information.
class MockNetworkInfo extends _i1.Mock implements _i18.NetworkInfo {
  MockNetworkInfo() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i7.Future<bool> get isConnected => (super.noSuchMethod(
        Invocation.getter(#isConnected),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);
}

/// A class which mocks [InternetConnectionChecker].
///
/// See the documentation for Mockito's code generation for more information.
class MockInternetConnectionChecker extends _i1.Mock
    implements _i4.InternetConnectionChecker {
  MockInternetConnectionChecker() {
    _i1.throwOnMissingStub(this);
  }

  @override
  Duration get checkInterval => (super.noSuchMethod(
        Invocation.getter(#checkInterval),
        returnValue: _FakeDuration_3(
          this,
          Invocation.getter(#checkInterval),
        ),
      ) as Duration);

  @override
  Duration get checkTimeout => (super.noSuchMethod(
        Invocation.getter(#checkTimeout),
        returnValue: _FakeDuration_3(
          this,
          Invocation.getter(#checkTimeout),
        ),
      ) as Duration);

  @override
  List<_i4.AddressCheckOptions> get addresses => (super.noSuchMethod(
        Invocation.getter(#addresses),
        returnValue: <_i4.AddressCheckOptions>[],
      ) as List<_i4.AddressCheckOptions>);

  @override
  set addresses(List<_i4.AddressCheckOptions>? value) => super.noSuchMethod(
        Invocation.setter(
          #addresses,
          value,
        ),
        returnValueForMissingStub: null,
      );

  @override
  _i7.Future<bool> get hasConnection => (super.noSuchMethod(
        Invocation.getter(#hasConnection),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<_i4.InternetConnectionStatus> get connectionStatus =>
      (super.noSuchMethod(
        Invocation.getter(#connectionStatus),
        returnValue: _i7.Future<_i4.InternetConnectionStatus>.value(
            _i4.InternetConnectionStatus.connected),
      ) as _i7.Future<_i4.InternetConnectionStatus>);

  @override
  _i7.Stream<_i4.InternetConnectionStatus> get onStatusChange =>
      (super.noSuchMethod(
        Invocation.getter(#onStatusChange),
        returnValue: _i7.Stream<_i4.InternetConnectionStatus>.empty(),
      ) as _i7.Stream<_i4.InternetConnectionStatus>);

  @override
  bool get hasListeners => (super.noSuchMethod(
        Invocation.getter(#hasListeners),
        returnValue: false,
      ) as bool);

  @override
  bool get isActivelyChecking => (super.noSuchMethod(
        Invocation.getter(#isActivelyChecking),
        returnValue: false,
      ) as bool);

  @override
  _i7.Future<_i4.AddressCheckResult> isHostReachable(
          _i4.AddressCheckOptions? options) =>
      (super.noSuchMethod(
        Invocation.method(
          #isHostReachable,
          [options],
        ),
        returnValue:
            _i7.Future<_i4.AddressCheckResult>.value(_FakeAddressCheckResult_4(
          this,
          Invocation.method(
            #isHostReachable,
            [options],
          ),
        )),
      ) as _i7.Future<_i4.AddressCheckResult>);
}

/// A class which mocks [SharedPreferences].
///
/// See the documentation for Mockito's code generation for more information.
class MockSharedPreferences extends _i1.Mock implements _i19.SharedPreferences {
  MockSharedPreferences() {
    _i1.throwOnMissingStub(this);
  }

  @override
  Set<String> getKeys() => (super.noSuchMethod(
        Invocation.method(
          #getKeys,
          [],
        ),
        returnValue: <String>{},
      ) as Set<String>);

  @override
  Object? get(String? key) => (super.noSuchMethod(Invocation.method(
        #get,
        [key],
      )) as Object?);

  @override
  bool? getBool(String? key) => (super.noSuchMethod(Invocation.method(
        #getBool,
        [key],
      )) as bool?);

  @override
  int? getInt(String? key) => (super.noSuchMethod(Invocation.method(
        #getInt,
        [key],
      )) as int?);

  @override
  double? getDouble(String? key) => (super.noSuchMethod(Invocation.method(
        #getDouble,
        [key],
      )) as double?);

  @override
  String? getString(String? key) => (super.noSuchMethod(Invocation.method(
        #getString,
        [key],
      )) as String?);

  @override
  bool containsKey(String? key) => (super.noSuchMethod(
        Invocation.method(
          #containsKey,
          [key],
        ),
        returnValue: false,
      ) as bool);

  @override
  List<String>? getStringList(String? key) =>
      (super.noSuchMethod(Invocation.method(
        #getStringList,
        [key],
      )) as List<String>?);

  @override
  _i7.Future<bool> setBool(
    String? key,
    bool? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setBool,
          [
            key,
            value,
          ],
        ),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<bool> setInt(
    String? key,
    int? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setInt,
          [
            key,
            value,
          ],
        ),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<bool> setDouble(
    String? key,
    double? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setDouble,
          [
            key,
            value,
          ],
        ),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<bool> setString(
    String? key,
    String? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setString,
          [
            key,
            value,
          ],
        ),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<bool> setStringList(
    String? key,
    List<String>? value,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #setStringList,
          [
            key,
            value,
          ],
        ),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<bool> remove(String? key) => (super.noSuchMethod(
        Invocation.method(
          #remove,
          [key],
        ),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<bool> commit() => (super.noSuchMethod(
        Invocation.method(
          #commit,
          [],
        ),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<bool> clear() => (super.noSuchMethod(
        Invocation.method(
          #clear,
          [],
        ),
        returnValue: _i7.Future<bool>.value(false),
      ) as _i7.Future<bool>);

  @override
  _i7.Future<void> reload() => (super.noSuchMethod(
        Invocation.method(
          #reload,
          [],
        ),
        returnValue: _i7.Future<void>.value(),
        returnValueForMissingStub: _i7.Future<void>.value(),
      ) as _i7.Future<void>);
}

/// A class which mocks [CustomHttpClient].
///
/// See the documentation for Mockito's code generation for more information.
class MockCustomHttpClient extends _i1.Mock implements _i18.CustomHttpClient {
  MockCustomHttpClient() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i7.Future<_i5.Response> get(String? endpoint) => (super.noSuchMethod(
        Invocation.method(
          #get,
          [endpoint],
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #get,
            [endpoint],
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.Response> post(
    String? endpoint, {
    Object? body,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #post,
          [endpoint],
          {#body: body},
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #post,
            [endpoint],
            {#body: body},
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.Response> put(
    String? endpoint, {
    Object? body,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #put,
          [endpoint],
          {#body: body},
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #put,
            [endpoint],
            {#body: body},
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.Response> delete(String? endpoint) => (super.noSuchMethod(
        Invocation.method(
          #delete,
          [endpoint],
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #delete,
            [endpoint],
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.StreamedResponse> send(_i5.BaseRequest? request) =>
      (super.noSuchMethod(
        Invocation.method(
          #send,
          [request],
        ),
        returnValue:
            _i7.Future<_i5.StreamedResponse>.value(_FakeStreamedResponse_6(
          this,
          Invocation.method(
            #send,
            [request],
          ),
        )),
      ) as _i7.Future<_i5.StreamedResponse>);
}

/// A class which mocks [CaseRepository].
///
/// See the documentation for Mockito's code generation for more information.
class MockCaseRepository extends _i1.Mock implements _i20.CaseRepository {
  MockCaseRepository() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>> createCase(
          _i21.CaseEntity? caseEntity) =>
      (super.noSuchMethod(
        Invocation.method(
          #createCase,
          [caseEntity],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>>.value(
            _FakeEither_0<_i8.Failure, _i21.CaseEntity>(
          this,
          Invocation.method(
            #createCase,
            [caseEntity],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>> updateCase(String? id) =>
      (super.noSuchMethod(
        Invocation.method(
          #updateCase,
          [id],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>>.value(
            _FakeEither_0<_i8.Failure, _i21.CaseEntity>(
          this,
          Invocation.method(
            #updateCase,
            [id],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>> deleteCase(String? id) =>
      (super.noSuchMethod(
        Invocation.method(
          #deleteCase,
          [id],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>.value(
            _FakeEither_0<_i8.Failure, _i2.Unit>(
          this,
          Invocation.method(
            #deleteCase,
            [id],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i2.Unit>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>> getCase(
          _i21.CaseEntity? caseEntity) =>
      (super.noSuchMethod(
        Invocation.method(
          #getCase,
          [caseEntity],
        ),
        returnValue: _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>>.value(
            _FakeEither_0<_i8.Failure, _i21.CaseEntity>(
          this,
          Invocation.method(
            #getCase,
            [caseEntity],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, _i21.CaseEntity>>);

  @override
  _i7.Future<_i2.Either<_i8.Failure, List<_i21.CaseEntity>>> getCases() =>
      (super.noSuchMethod(
        Invocation.method(
          #getCases,
          [],
        ),
        returnValue:
            _i7.Future<_i2.Either<_i8.Failure, List<_i21.CaseEntity>>>.value(
                _FakeEither_0<_i8.Failure, List<_i21.CaseEntity>>(
          this,
          Invocation.method(
            #getCases,
            [],
          ),
        )),
      ) as _i7.Future<_i2.Either<_i8.Failure, List<_i21.CaseEntity>>>);
}

/// A class which mocks [Client].
///
/// See the documentation for Mockito's code generation for more information.
class MockHttpClient extends _i1.Mock implements _i5.Client {
  MockHttpClient() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i7.Future<_i5.Response> head(
    Uri? url, {
    Map<String, String>? headers,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #head,
          [url],
          {#headers: headers},
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #head,
            [url],
            {#headers: headers},
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.Response> get(
    Uri? url, {
    Map<String, String>? headers,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #get,
          [url],
          {#headers: headers},
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #get,
            [url],
            {#headers: headers},
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.Response> post(
    Uri? url, {
    Map<String, String>? headers,
    Object? body,
    _i22.Encoding? encoding,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #post,
          [url],
          {
            #headers: headers,
            #body: body,
            #encoding: encoding,
          },
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #post,
            [url],
            {
              #headers: headers,
              #body: body,
              #encoding: encoding,
            },
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.Response> put(
    Uri? url, {
    Map<String, String>? headers,
    Object? body,
    _i22.Encoding? encoding,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #put,
          [url],
          {
            #headers: headers,
            #body: body,
            #encoding: encoding,
          },
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #put,
            [url],
            {
              #headers: headers,
              #body: body,
              #encoding: encoding,
            },
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.Response> patch(
    Uri? url, {
    Map<String, String>? headers,
    Object? body,
    _i22.Encoding? encoding,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #patch,
          [url],
          {
            #headers: headers,
            #body: body,
            #encoding: encoding,
          },
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #patch,
            [url],
            {
              #headers: headers,
              #body: body,
              #encoding: encoding,
            },
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<_i5.Response> delete(
    Uri? url, {
    Map<String, String>? headers,
    Object? body,
    _i22.Encoding? encoding,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #delete,
          [url],
          {
            #headers: headers,
            #body: body,
            #encoding: encoding,
          },
        ),
        returnValue: _i7.Future<_i5.Response>.value(_FakeResponse_5(
          this,
          Invocation.method(
            #delete,
            [url],
            {
              #headers: headers,
              #body: body,
              #encoding: encoding,
            },
          ),
        )),
      ) as _i7.Future<_i5.Response>);

  @override
  _i7.Future<String> read(
    Uri? url, {
    Map<String, String>? headers,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #read,
          [url],
          {#headers: headers},
        ),
        returnValue: _i7.Future<String>.value(_i17.dummyValue<String>(
          this,
          Invocation.method(
            #read,
            [url],
            {#headers: headers},
          ),
        )),
      ) as _i7.Future<String>);

  @override
  _i7.Future<_i23.Uint8List> readBytes(
    Uri? url, {
    Map<String, String>? headers,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #readBytes,
          [url],
          {#headers: headers},
        ),
        returnValue: _i7.Future<_i23.Uint8List>.value(_i23.Uint8List(0)),
      ) as _i7.Future<_i23.Uint8List>);

  @override
  _i7.Future<_i5.StreamedResponse> send(_i5.BaseRequest? request) =>
      (super.noSuchMethod(
        Invocation.method(
          #send,
          [request],
        ),
        returnValue:
            _i7.Future<_i5.StreamedResponse>.value(_FakeStreamedResponse_6(
          this,
          Invocation.method(
            #send,
            [request],
          ),
        )),
      ) as _i7.Future<_i5.StreamedResponse>);

  @override
  void close() => super.noSuchMethod(
        Invocation.method(
          #close,
          [],
        ),
        returnValueForMissingStub: null,
      );
}

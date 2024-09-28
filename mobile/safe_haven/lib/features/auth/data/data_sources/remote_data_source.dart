import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;
import 'package:safe_haven/core/constants/constants.dart';
import 'package:safe_haven/core/error/exception.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/data/models/authenticated_model.dart';
import 'package:safe_haven/features/auth/data/models/log_in_model.dart';
import 'package:safe_haven/features/auth/data/models/reset_password_model.dart';
import 'package:safe_haven/features/auth/data/models/sign_up_model.dart';

abstract class AuthenticationRemoteDataSource {
  /// calls the http://loginendpoint
  Future<AuthenticatedModel> login(LogInModel logInModel);

  /// calls the http://signupendpoint
  Future<Unit> signup(SignUpModel signUpModel);

  /// calls the http://forgot-endpoint
  Future<Unit> forgotPassword(String resetEmail);

  /// calls the http://reset-endpoint
  Future<Unit> resetPassword(ResetPasswordModel resetPasswordModel);

  /// calls the http://reset-endpoint
  Future<Unit> googleLogin();
}

class AuthRemoteDataSourceImpl extends AuthenticationRemoteDataSource {
  final http.Client client;

  AuthRemoteDataSourceImpl({required this.client});

  @override
  Future<AuthenticatedModel> login(LogInModel logInModel) async {
    var uri = Uri.parse('${Urls.authUrl}/login');
    print(logInModel.toJson());
    print('exi dersual atleast in the login api');

    try {
      final response = await client.post(uri,
          body: jsonEncode(
            logInModel.toJson(),
          ),
          headers: {'Content-Type': 'application/json'});
      print(response.statusCode);
      print('ezi');
      print(response.body);
      if (response.statusCode == 200) {
        print(
          'lets print first',
        );
        print(AuthenticatedModel.fromJson(json.decode(response.body)['data']));
        return AuthenticatedModel.fromJson(json.decode(response.body)['data']);
      } else if (response.statusCode == 500) {
        print('There was no email associated with this account');
      }
      print(
          'this is the test version in the remote_data_source(login) where u try to check the custom status of samri\'s');
      throw const ServerException(ErrorMessages.serverError);
    } on SocketException {
      throw SocketException(ErrorMessages.socketError);
    } on NotFoundException {
      throw NotFoundException(ErrorMessages.notFoundError);
    }
  }

  @override
  Future<Unit> signup(SignUpModel signUpModel) async {
    var uri = Uri.parse('${Urls.authUrl}/register');
    print('ezi dersual in the signup api');
    print(signUpModel.toJson());
    try {
      final response = await client
          .post(uri, body: jsonEncode(signUpModel.toJson()), headers: {
        'Content-Type': 'application/json',
      });
      print(response.body);
      if (response.statusCode == 201) {
        return unit;
      } else if (response.statusCode == 500) {
        throw UserAlreadyExistsException();
      } else {
        throw ServerException('server error in signing up');
      }
    } on SocketException {
      throw SocketException(ErrorMessages.socketError);
    }
  }

  @override
  Future<Unit> forgotPassword(String resetEmail) async {
    var uri = Uri.parse('${Urls.authUrl}/forgot-password');
    print('ezi dersual in the forgot api');
    try {
      print(jsonEncode(resetEmail));
      final response =
          await client.post(uri, body: jsonEncode(resetEmail), headers: {
        'Content-Type': 'application/json',
      });
      print(response.statusCode);
      print(response.body);
      if (response.statusCode == 200) {
        return unit;
      } else if (response.statusCode == 500) {
        throw NotFoundException(ErrorMessages.notFoundError);
      } else {
        throw ServerException('server error in sending email to reset email');
      }
    } on SocketException {
      throw SocketException(ErrorMessages.socketError);
    }
  }

  static const token =
      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjo0ODc2MCwiZW1haWwiOiJMaXlhLmRhbmllbC56ZWxla2VAZ21haWwuY29tIiwiZXhwIjoxNzI3MDg4MzEyfQ.yVnejOPdAp2RtmTLth5yTjoZdlmGBYcrWp3DrRe3_mI';
  @override
  Future<Unit> resetPassword(ResetPasswordModel resetPasswordModel) async {
    var uri = Uri.parse('${Urls.authUrl}/reset-password');
    // print('ezi dersual in the resetpassword api');
    // print(resetPasswordModel.toJson());
    try {
      final response = await client
          .post(uri, body: jsonEncode(resetPasswordModel.toJson()), headers: {
        'Content-Type': 'application/json',
      });
      print(response.body);
      if (response.statusCode == 201) {
        return unit;
      } else if (response.statusCode == 500) {
        throw UserAlreadyExistsException();
      } else {
        throw const ServerException('server error in signing up');
      }
    } on SocketException {
      throw SocketException(ErrorMessages.socketError);
    }
  }

  @override
  Future<Unit> googleLogin() async {
    var uri = Uri.parse('${Urls.authUrl}/auth/google');
    print('ezi dersual in the google sign in api');

    try {
      final response = await client.get(uri);

      // if (response.statusCode == 201) {
      print(response);
      return unit;
      // } else if (response.statusCode == 500) {
      //   throw UserAlreadyExistsException();
      // } else {
      //   throw const ServerException('server error in signing up');
      // }
    } on ServerException {
      throw const ServerException('server error in googls sign in');
    }
  }

  // @override
  // Future<UserDataModel> getuser(String token) async{
  //   var uri = Uri.parse('${Urls.authUrl}/users/me');
  //   try{
  //     final response = await client.get(uri, headers: {'Authorization' : 'Bearer $token'});
  //     if (response.statusCode == 200){
  //       return UserDataModel.fromJson(json.decode(response.body)['data']);
  //     } else if (response.statusCode == 401){
  //       throw UnauthorizedException();
  //     } else{
  //       throw ServerException();
  //     }
  //   } on SocketException{
  //     throw const SocketException(ErrorMessages.socketError);
  //   }
  // }

  // @override
  // Future<AuthenticatedModel> login(LoginModel login_model) async{
  //   var uri = Uri.parse('${Urls.authUrl}/auth/login');
  //   try{
  //     final response = await client.post(uri, body: login_model.toJson() );
  //     if (response.statusCode == 201){
  //       return AuthenticatedModel.fromJson(json.decode(response.body)['data']);
  //     } else if (response.statusCode == 401) {
  //       throw UnauthorizedException();
  //     }else{
  //       throw ServerException();
  //     }
  //   } on SocketException{
  //     throw const SocketException(ErrorMessages.socketError);
  //   }
  // }

  // @override
  // Future<Unit> register(RegisterModel register_model) async{
  //   var uri = Uri.parse('${Urls.authUrl}/auth/register');

  //   try{
  //     final response = await client.post(uri, body: register_model.toJson());
  //     if (response.statusCode == 201){
  //       return unit;
  //     } else if(response.statusCode == 409){
  //       throw UserAlreadyExistsException();
  //     } else{
  //       throw ServerException();
  //     }
  //   } on SocketException{
  //     throw const SocketException(ErrorMessages.socketError);
  //   }
  // }
}

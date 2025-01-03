import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:safe_haven/core/error/exception.dart';
import 'package:shared_preferences/shared_preferences.dart';

abstract class AuthenticationLocalDataSource {
  Future<String> getToken();
  Future<Unit> cacheTokens(String token, String refreshToken);
  Future<Unit> logout();
  Future<String> getRefreshToken();
}

const TOKEN = 'token';
const REFRESH_TOKEN = 'refreshToken';

class AuthLocalDataSourceImpl extends AuthenticationLocalDataSource {
  final SharedPreferences sharedPreferences;

  AuthLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<Unit> cacheTokens(String token, String refreshToken) {
    try {
      final jsonToken = json.encode(token);
      final jsonRefreshToken = json.encode(refreshToken);
      sharedPreferences.setString(TOKEN, jsonToken);
      sharedPreferences.setString(REFRESH_TOKEN, jsonRefreshToken);
      return Future.value(unit);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<String> getToken() {
    try {
      final token = sharedPreferences.getString(TOKEN);
      if (token != null) {
        final decodedToken = json.decode((token));
        return Future.value(decodedToken);
      } else {
        throw CacheException();
      }
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<String> getRefreshToken() {
    try {
      final refreshToken = sharedPreferences.getString(REFRESH_TOKEN);
      if (refreshToken != null) {
        final decodedRefreshToken = json.decode(refreshToken);
        return Future.value(decodedRefreshToken);
      } else {
        throw CacheException();
      }
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<Unit> logout() {
    try {
      sharedPreferences.remove(TOKEN);
      return Future.value(unit);
    } catch (e) {
      throw CacheException();
    }
  }
}

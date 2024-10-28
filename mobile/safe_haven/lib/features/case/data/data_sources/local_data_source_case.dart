import 'dart:convert';

import 'package:safe_haven/features/case/data/models/case_model.dart';
import 'package:shared_preferences/shared_preferences.dart';

abstract class CaseLocalDataSource {
  //gets all cases stored last time
  Future<List<CaseModel>> getAllCases();

  //gets all cases stored last time
  Future<List<CaseModel>> getCase();

  //caches Last Product
  Future<void> cacheCase(CaseModel caseModel);

  Future<void> cacheLastProducts(List<CaseModel> caseModels);
  
  
}

const CACHED_CASE = 'CACHED_CASE';
const CACHED_CASES = 'CACHED_CASES';


class CaseLocalDataSourceImpl implements CaseLocalDataSource {
  final SharedPreferences sharedPreferences;

  CaseLocalDataSourceImpl({required this.sharedPreferences});
  @override
  Future<void> cacheCase(CaseModel caseModel) {
    return sharedPreferences.setString(
        CACHED_CASE, json.encode(caseModel.toJson()));
  }

  @override
  Future<void> cacheLastProducts(List<CaseModel> caseModels) {
    return sharedPreferences.setString(
        CACHED_CASES, convertToJsonList(caseModels));
  }

  @override
  Future<List<CaseModel>> getAllCases() {
    // TODO: implement getAllCases
    throw UnimplementedError();
  }

  @override
  Future<List<CaseModel>> getCase() {
    // TODO: implement getCase
    throw UnimplementedError();
  }
}

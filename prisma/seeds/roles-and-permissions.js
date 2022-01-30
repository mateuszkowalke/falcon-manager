"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
exports.__esModule = true;
exports.seedRolesAndPermissions = void 0;
var client_1 = require("@prisma/client");
var RolesEnum;
(function (RolesEnum) {
    RolesEnum["SUPER_ADMIN"] = "Super Admin";
    RolesEnum["BREEDER"] = "Breeder";
    RolesEnum["WORKER"] = "Worker";
})(RolesEnum || (RolesEnum = {}));
var PermissionEnum;
(function (PermissionEnum) {
    PermissionEnum["GetUsers"] = "GET_USERS";
    PermissionEnum["CreateUser"] = "CREATE_USER";
    PermissionEnum["UpdateUser"] = "UPDATE_USER";
    PermissionEnum["DeleteUser"] = "DELETE_USER";
    PermissionEnum["GetBreedingProjects"] = "GET_BREEDING_PROJECTS";
    PermissionEnum["CreateBreedingProject"] = "CREATE_BREEDING_PROJECT";
    PermissionEnum["UpdateBreedingProject"] = "UPDATE_BREEDING_PROJECT";
    PermissionEnum["DeleteBreedingProject"] = "DELETE_BREEDING_PROJECT";
    PermissionEnum["GetFalcons"] = "GET_FALCONS";
    PermissionEnum["CreateFalcon"] = "CREATE_FALCON";
    PermissionEnum["UpdateFalcon"] = "UPDATE_FALCON";
    PermissionEnum["DeleteFalcon"] = "DELETE_FALCON";
})(PermissionEnum || (PermissionEnum = {}));
var prisma = new client_1.PrismaClient();
function seedRolesAndPermissions() {
    return __awaiter(this, void 0, void 0, function () {
        var adminRole, breederRole, workerRole, getUsersPerm, createUserPerm, updateUserPerm, deleteUserPerm, getBreedingProjectsPerm, createBreedingProjectPerm, updateBreedingProjectPerm, deleteBreedingProjectPerm, getFalconsPerm, createFalconPerm, updateFalconPerm, deleteFalconPerm;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, createRoleIfNotExist(RolesEnum.SUPER_ADMIN)];
                case 1:
                    adminRole = _a.sent();
                    return [4 /*yield*/, createRoleIfNotExist(RolesEnum.BREEDER)];
                case 2:
                    breederRole = _a.sent();
                    return [4 /*yield*/, createRoleIfNotExist(RolesEnum.WORKER)];
                case 3:
                    workerRole = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.GetUsers)];
                case 4:
                    getUsersPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.CreateUser)];
                case 5:
                    createUserPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.UpdateUser)];
                case 6:
                    updateUserPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.DeleteUser)];
                case 7:
                    deleteUserPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.GetBreedingProjects)];
                case 8:
                    getBreedingProjectsPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.CreateBreedingProject)];
                case 9:
                    createBreedingProjectPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.UpdateBreedingProject)];
                case 10:
                    updateBreedingProjectPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.DeleteBreedingProject)];
                case 11:
                    deleteBreedingProjectPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.GetFalcons)];
                case 12:
                    getFalconsPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.CreateFalcon)];
                case 13:
                    createFalconPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.UpdateFalcon)];
                case 14:
                    updateFalconPerm = _a.sent();
                    return [4 /*yield*/, createPermissionIfNotExist(PermissionEnum.DeleteFalcon)];
                case 15:
                    deleteFalconPerm = _a.sent();
                    // connect permissions with role
                    // Super Admin
                    return [4 /*yield*/, connectRoleWithPermissions(adminRole, getUsersPerm, createUserPerm, updateUserPerm, deleteUserPerm, getBreedingProjectsPerm, createBreedingProjectPerm, updateBreedingProjectPerm, deleteBreedingProjectPerm, getFalconsPerm, createFalconPerm, updateFalconPerm, deleteFalconPerm)];
                case 16:
                    // connect permissions with role
                    // Super Admin
                    _a.sent();
                    // Breeder
                    return [4 /*yield*/, connectRoleWithPermissions(breederRole, getUsersPerm, createUserPerm, updateUserPerm, deleteUserPerm, getBreedingProjectsPerm, createBreedingProjectPerm, updateBreedingProjectPerm, deleteBreedingProjectPerm, getFalconsPerm, createFalconPerm, updateFalconPerm, deleteFalconPerm)];
                case 17:
                    // Breeder
                    _a.sent();
                    // Worker
                    return [4 /*yield*/, connectRoleWithPermissions(workerRole, getUsersPerm, createUserPerm, updateUserPerm, deleteUserPerm, getBreedingProjectsPerm, getFalconsPerm)];
                case 18:
                    // Worker
                    _a.sent();
                    return [2 /*return*/];
            }
        });
    });
}
exports.seedRolesAndPermissions = seedRolesAndPermissions;
function createRoleIfNotExist(name) {
    return __awaiter(this, void 0, void 0, function () {
        var role;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, prisma.role.findFirst({
                        where: { name: name },
                        include: { permissions: true }
                    })];
                case 1:
                    role = _a.sent();
                    if (!!role) return [3 /*break*/, 3];
                    return [4 /*yield*/, prisma.role.create({
                            data: { name: name },
                            include: { permissions: true }
                        })];
                case 2:
                    role = _a.sent();
                    _a.label = 3;
                case 3: return [2 /*return*/, role];
            }
        });
    });
}
function createPermissionIfNotExist(name) {
    return __awaiter(this, void 0, void 0, function () {
        var permission;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, prisma.permission.findFirst({
                        where: { name: name },
                        include: { roles: true }
                    })];
                case 1:
                    permission = _a.sent();
                    if (!!permission) return [3 /*break*/, 3];
                    return [4 /*yield*/, prisma.permission.create({
                            data: { name: name },
                            include: { roles: true }
                        })];
                case 2:
                    permission = _a.sent();
                    _a.label = 3;
                case 3: return [2 /*return*/, permission];
            }
        });
    });
}
function connectRoleWithPermissions(role) {
    var permissions = [];
    for (var _i = 1; _i < arguments.length; _i++) {
        permissions[_i - 1] = arguments[_i];
    }
    return __awaiter(this, void 0, void 0, function () {
        var _a, permissions_1, permission;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    _a = 0, permissions_1 = permissions;
                    _b.label = 1;
                case 1:
                    if (!(_a < permissions_1.length)) return [3 /*break*/, 4];
                    permission = permissions_1[_a];
                    if (!!role.permissions.map(function (_a) {
                        var permissionId = _a.permissionId;
                        return permissionId;
                    }).includes(permission.id)) return [3 /*break*/, 3];
                    return [4 /*yield*/, prisma.rolesOnPermissions.create({
                            data: {
                                roleId: role.id,
                                permissionId: permission.id
                            }
                        })];
                case 2:
                    _b.sent();
                    _b.label = 3;
                case 3:
                    _a++;
                    return [3 /*break*/, 1];
                case 4: return [2 /*return*/];
            }
        });
    });
}

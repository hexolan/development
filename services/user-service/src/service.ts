import { Types } from "mongoose";
import { ConnectError, Code } from "@connectrpc/connect";

import { User, IUser } from "./mongo/User";
import userProducer from "./kafka/producer";

function isValidUsername(username: string): boolean {
  let length = username.length;
  if (length < 3 || length > 32) {
    return false
  }

  const regexCheck = new RegExp("^[^_]\\w+[^_]$");
  if (!regexCheck.test(username)) {
    return false
  }

  return true
}

async function createUser(username: string): Promise<IUser> {
  if (!isValidUsername(username)) {
    throw new ConnectError("invalid username", Code.InvalidArgument);
  }

  let newUser = new User({ username: username });

  let user = await newUser.save().then(async (user) => {
    await userProducer.sendCreatedEvent(user);
    return user;
  }).catch((error) => {
    // todo: ensure error is a result of unique constraint violation
    throw new ConnectError("username already exists", Code.AlreadyExists);
  });

  return user;
}

async function getUserById(id: string): Promise<IUser> {
  // ensure id is valid
  if (!Types.ObjectId.isValid(id)) {
    throw new ConnectError("invalid id provided", Code.InvalidArgument);
  }

  // attempt to get the user document
  let user = await User.findById(id).exec()
  if (user === null) {
    throw new ConnectError("user not found", Code.NotFound);
  }

  return user
}

async function getUserByUsername(username: string): Promise<IUser> {
  // ensure username is valid
  if (username === "") {
    throw new ConnectError("invalid username", Code.InvalidArgument)
  }

  // attempt to find the document
  let user = await User.findOne({ username: username })
  if (user === null) {
    throw new ConnectError("user not found", Code.NotFound);
  }

  return user
}

async function updateUserById(id: string, newUsername: string): Promise<IUser> {
  if (!isValidUsername(newUsername)) {
    throw new ConnectError("invalid username value", Code.InvalidArgument);
  }

  // ensure id is valid
  if (!Types.ObjectId.isValid(id)) {
    throw new ConnectError("invalid id provided", Code.InvalidArgument)
  }

  // attempt to update the user
  const updatedUser = await User.findByIdAndUpdate(
    id,
    { username: newUsername },
    { new: true }
  ).catch(
    (error) => {
      // todo: check unique constraint violation
      throw new ConnectError("user not found", Code.NotFound);
    }
  )

  if (updatedUser === null) {
    throw new ConnectError("something unexpected went wrong", Code.Internal);
  }

  return updatedUser;
}

async function updateUserByUsername(username: string, newUsername: string): Promise<IUser> {
  if (!isValidUsername(newUsername)) {
    throw new ConnectError("invalid username value", Code.InvalidArgument);
  }

  // attempt to update the user
  const updatedUser = await User.findOneAndUpdate(
    { username: username },
    { username: newUsername },
    { new: true }
  ).catch(
    (error) => {
      // todo: catch username not unique error
      throw new ConnectError("user not found", Code.NotFound);
    }
  )

  if (updatedUser === null || updatedUser === undefined) {
    throw new ConnectError("something unexpected went wrong", Code.Internal);
  }

  return updatedUser;
}

async function deleteUserById(id: string): Promise<void> {
  // ensure id is valid
  if (!Types.ObjectId.isValid(id)) {
    throw new ConnectError("invalid id provided", Code.InvalidArgument)
  }

  // atempt to delete the user
  await User.findByIdAndDelete(id).catch(() => {
    throw new ConnectError("user not found", Code.NotFound)
  });
}

async function deleteUserByUsername(username: string): Promise<void> {
  // attempt to delete the user
  await User.findOneAndDelete({ username: username }).catch(() => {
    throw new ConnectError("user not found", Code.NotFound)
  });
}

export { createUser, getUserById, getUserByUsername, updateUserById, updateUserByUsername, deleteUserById, deleteUserByUsername }
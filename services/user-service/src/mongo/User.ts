import { Document, Schema, model } from "mongoose";
import uniqueValidator from "mongoose-unique-validator";

const userSchema = new Schema(
  {
    username: { type: String, required: true, lowercase: true, unique: true }
  },
  { 
    timestamps: true
  }
);

userSchema.plugin(uniqueValidator);

interface IUser extends Document {
  username: string;
  createdAt: Date;
  updatedAt: Date;
}

const User = model<IUser>("User", userSchema);

export { User, IUser }
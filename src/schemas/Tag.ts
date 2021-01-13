import { model, Schema, Document } from 'mongoose';

type Tag = Document & {};

const TagSchema = new Schema(
  {
    title: {
      type: String,
      lowercase: true,
      trim: true,
      unique: true,
      required: true,
    },
    color: {
      type: String,
      trim: true,
    },
  },
  {
    timestamps: true,
  },
);

export default model<Tag>('tags', TagSchema);

import { model, Schema, Document } from 'mongoose';

type Contact = Document & {};

const ContactSchema = new Schema(
  {
    email: {
      type: String,
      lowercase: true,
      trim: true,
      unique: true,
      required: true,
    },
    tags: [
      {
        type: Schema.Types.ObjectId,
        ref: 'tags',
      },
    ],
  },
  {
    timestamps: true,
  },
);

export default model<Contact>('contacts', ContactSchema);

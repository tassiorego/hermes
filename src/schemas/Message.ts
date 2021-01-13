import { model, Schema, Document } from 'mongoose';

type Message = Document & {};

const MessageSchema = new Schema(
  {
    subject: {
      type: String,
      trim: true,
      required: true,
    },
    body: {
      type: String,
      required: true,
    },
    completedAt: {
      type: Date,
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

export default model<Message>('messages', MessageSchema);

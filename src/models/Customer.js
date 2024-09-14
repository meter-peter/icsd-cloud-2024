const mongoose = require('mongoose');
const { v4: uuidv4 } = require('uuid');

const customerSchema = new mongoose.Schema({
  _id: { type: String, default: uuidv4 },
  createdAt: { type: Date, default: Date.now },
  identityNumber: {
    type: String,
    required: true,
    unique: true,
    match: /^[A-Z]{2}\d{6}$/,
  },
  firstName: {
    type: String,
    required: true,
    match: /^[A-Z][a-z]{2,}$/,
  },
  lastName: {
    type: String,
    required: true,
    match: /^[A-Z][a-z]{2,}$/,
  },
  gender: {
    type: String,
    enum: ['male', 'female'],
    required: true,
  },
  birthDate: {
    type: Date,
    required: true,
    validate: {
      validator: function(date) {
        return (new Date().getFullYear() - date.getFullYear()) >= 16;
      },
      message: 'Customer must be at least 16 years old',
    },
  },
  addresses: [{
    type: String,
    validate: {
      validator: function(v) {
        return /^[\w\s]+\s\d{1,3}$/.test(v);
      },
      message: props => `${props.value} is not a valid address`,
    },
  }],
  phoneNumbers: [{
    type: String,
    validate: {
      validator: function(v) {
        return /^\d{10}$/.test(v);
      },
      message: props => `${props.value} is not a valid phone number`,
    },
  }],
});

module.exports = mongoose.model('Customer', customerSchema);
const express = require('express');
const customerController = require('../controllers/customerController');

const router = express.Router();

router.post('/', customerController.createCustomer);
router.get('/:id', customerController.getCustomer);
router.get('/', customerController.getAllCustomers);
router.get('/search', customerController.searchCustomers);
router.delete('/:id', customerController.deleteCustomer);
router.put('/:id', customerController.updateCustomer);

module.exports = router;


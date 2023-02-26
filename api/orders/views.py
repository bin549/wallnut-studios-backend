from flask_restx import Namespace, Resource

order_namespace = Namespace('orders', description="Namespace for orders")


@order_namespace.route('/orders/')
class OrderGetCreate(Resource):
    def get(self):
        pass

    def post(self):
        pass


@order_namespace.route('/order/<int:order_id>')
class GetUpdateDelete(Resource):
    def get(self):
        pass

    def put(self):
        pass

    def delete(self):
        pass


@order_namespace.route('/user/<int:user_id>/order/<int:order_id>/')
class GetSpecificOrderByUser(Resource):
    def get(self):
        pass


@order_namespace.route('/order/status/<int:order_id>')
class UpdateOrderStatus(Resource):
    def patch(self, order_id):
        pass

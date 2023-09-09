import base64

from auth_service.models.proto import user_pb2

thing = user_pb2.UserEvent(
    type="deleted",
    data=user_pb2.User(
        id="abc",
        username="test"
    )
)
print(thing.SerializeToString())
print("NEWLINE\n" + thing.SerializeToString().decode("utf-8") + "\nENDLINE")
print(base64.standard_b64encode(thing.SerializeToString()))
namespace TyphenApi.Type.Submarine.Battle
{
    [MessagePack.MessagePackObject]
    [Newtonsoft.Json.JsonObject(Newtonsoft.Json.MemberSerialization.OptIn)]
    public partial class Ping : TyphenApi.TypeBase<Ping>, TyphenApi.IRealTimeMessage
    {
        [TyphenApi.QueryStringProperty("message", false)]
        [MessagePack.Key("message")]
        [Newtonsoft.Json.JsonProperty("message")]
        [Newtonsoft.Json.JsonRequired]
        public string Message { get; set; }
    }
}

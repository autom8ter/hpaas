<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>meshpaas.Networking</code>
 */
class Networking extends \Google\Protobuf\Internal\Message
{
    /**
     * gateways to bind to
     *
     * Generated from protobuf field <code>repeated string gateways = 1;</code>
     */
    private $gateways;
    /**
     * host names to bind to
     *
     * Generated from protobuf field <code>repeated string hosts = 2;</code>
     */
    private $hosts;
    /**
     * export service to other applications in other namespaces
     *
     * Generated from protobuf field <code>bool export = 3;</code>
     */
    private $export = false;
    /**
     * http route matchers/configurations
     *
     * Generated from protobuf field <code>repeated .meshpaas.HTTPRoute http_routes = 4;</code>
     */
    private $http_routes;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $gateways
     *           gateways to bind to
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $hosts
     *           host names to bind to
     *     @type bool $export
     *           export service to other applications in other namespaces
     *     @type \Meshpaas\HTTPRoute[]|\Google\Protobuf\Internal\RepeatedField $http_routes
     *           http route matchers/configurations
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * gateways to bind to
     *
     * Generated from protobuf field <code>repeated string gateways = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getGateways()
    {
        return $this->gateways;
    }

    /**
     * gateways to bind to
     *
     * Generated from protobuf field <code>repeated string gateways = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setGateways($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->gateways = $arr;

        return $this;
    }

    /**
     * host names to bind to
     *
     * Generated from protobuf field <code>repeated string hosts = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getHosts()
    {
        return $this->hosts;
    }

    /**
     * host names to bind to
     *
     * Generated from protobuf field <code>repeated string hosts = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setHosts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->hosts = $arr;

        return $this;
    }

    /**
     * export service to other applications in other namespaces
     *
     * Generated from protobuf field <code>bool export = 3;</code>
     * @return bool
     */
    public function getExport()
    {
        return $this->export;
    }

    /**
     * export service to other applications in other namespaces
     *
     * Generated from protobuf field <code>bool export = 3;</code>
     * @param bool $var
     * @return $this
     */
    public function setExport($var)
    {
        GPBUtil::checkBool($var);
        $this->export = $var;

        return $this;
    }

    /**
     * http route matchers/configurations
     *
     * Generated from protobuf field <code>repeated .meshpaas.HTTPRoute http_routes = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getHttpRoutes()
    {
        return $this->http_routes;
    }

    /**
     * http route matchers/configurations
     *
     * Generated from protobuf field <code>repeated .meshpaas.HTTPRoute http_routes = 4;</code>
     * @param \Meshpaas\HTTPRoute[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setHttpRoutes($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Meshpaas\HTTPRoute::class);
        $this->http_routes = $arr;

        return $this;
    }

}

